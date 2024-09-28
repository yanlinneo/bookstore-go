package models

import (
	"bookstore-go/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                 int64  `json:"id"`
	Username           string `json:"username" binding:"required"`
	hashedPassword     string `json:"-"`
	RoleID             int64  `json:"role_id"`
	MustChangePassword bool   `json:"-"`
}

type LoginUser struct {
	Username           string `json:"username" binding:"required"`
	Password           string `json:"password" binding:"required"`
	NewPassword        string `json:"new_password"`
	MustChangePassword bool   `json:"-"`
}

type RegisterUser struct {
	Username           string `json:"username" binding:"required"`
	Password           string `json:"password" binding:"required"`
	RoleID             int64  `json:"-"`
	MustChangePassword bool   `json:"-"`
}

func (loginUser *LoginUser) Authenticate() (*User, error) {
	query := `
	SELECT id, password, must_change_password
	FROM users 
	WHERE username = ? 
	`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var user User
	row := stmt.QueryRow(&loginUser.Username)
	err = row.Scan(&user.ID, &user.hashedPassword, &loginUser.MustChangePassword)
	if err != nil {
		return nil, err
	}

	if !compare(user.hashedPassword, loginUser.Password) {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func (registerUser *RegisterUser) Create() error {
	query := `
	INSERT INTO users(username, password, role_id, must_change_password)
	VALUES(?, ?, ?, ?)
	`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	registerUser.Password, err = hash(registerUser.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(&registerUser.Username, &registerUser.Password, &registerUser.RoleID, &registerUser.MustChangePassword)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (user *User) UpdatePassword(loginUser *LoginUser) error {
	query := `
	UPDATE users
	SET password = ?, must_change_password = ?
	WHERE id = ?
	`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	user.hashedPassword, err = hash(loginUser.NewPassword)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.hashedPassword, 0, user.ID)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func find(username string) error {
	query := `
	SELECT id
	FROM users 
	WHERE username = ? 
	`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	var userID int64
	row := stmt.QueryRow(&username)
	err = row.Scan(&userID)
	if err != nil {
		return err // no rows = create admin
	}

	return nil
}

func hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}

func compare(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func CreateAdmin() error {
	username := "admin"

	if err := find(username); err == nil {
		return nil //admin exists. don't need to create
	}

	admin := RegisterUser{Username: username, Password: "hello_admin123", RoleID: 1, MustChangePassword: true}
	if err := admin.Create(); err != nil {
		return err
	}

	return nil
}
