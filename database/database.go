package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createBooksTable := `
	CREATE TABLE IF NOT EXISTS books(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		vbid VARCHAR(50) UNIQUE NOT NULL,
		title VARCHAR(300) NOT NULL,
		author VARCHAR(300) NOT NULL,
		user_id INTEGER NOT NULL,
		updated_at DATETIME NOT NULL
	)
	`

	_, err := DB.Exec(createBooksTable)
	if err != nil {
		panic("Books table could not be created.")
	}

	createRolesTable := `
	CREATE TABLE IF NOT EXISTS roles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name username varchar(100) UNIQUE NOT NULL
	);

	INSERT OR IGNORE INTO roles (name) VALUES ("admin"), ("manager");
	`

	_, err = DB.Exec(createRolesTable)
	if err != nil {
		panic(err)
	}

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username varchar(100) UNIQUE NOT NULL,
		password TEXT UNIQUE NOT NULL,
		role_id INTEGER NOT NULL,
		must_change_password TINYINT NOT NULL DEFAULT 1,
		FOREIGN KEY (role_id) REFERENCES roles(id)
	)
	`

	res, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Users table could not be created.")
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected > 0 {

	}
}
