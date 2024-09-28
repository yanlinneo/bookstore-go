package models

import (
	"bookstore-go/database"
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	VBID      string    `json:"vbid" binding:"required"`
	Title     string    `json:"title" binding:"required"`
	Author    string    `json:"author" binding:"required"`
	UserID    int64     `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at" binding:"required"`
}

func AllBooks() ([]Book, error) {
	query := `SELECT * FROM books`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book = []Book{}
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.VBID, &book.Title, &book.Author, &book.UserID, &book.UpdatedAt)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func BookByID(id int64) (*Book, error) {
	query := `
	SELECT * FROM books
	WHERE id = ?
	`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var book Book
	row.Scan(&book.ID, &book.VBID, &book.Title, &book.Author, &book.UserID, &book.UpdatedAt)

	if book.ID == 0 {
		return nil, nil
	}

	return &book, nil
}

func BookByVBID(vbid string) (*Book, error) {
	query := `
	SELECT * FROM books
	WHERE vbid = ?
	`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(vbid)

	var book Book
	row.Scan(&book.ID, &book.VBID, &book.Title, &book.Author, &book.UserID, &book.UpdatedAt)

	if book.ID == 0 {
		return nil, nil
	}

	return &book, nil
}

func (book *Book) Create() error {
	query := `
	INSERT INTO books(vbid, title, author, user_id, updated_at)
	VALUES(?, ?, ?, ?, ?)
	`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(book.VBID, book.Title, book.Author, book.UserID, book.UpdatedAt)
	if err != nil {
		return err
	}

	book.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (book *Book) Update() error {
	query := `
	UPDATE books
	SET vbid = ?, title = ?, author = ?, user_id = ?, updated_at = ?
	WHERE id = ?
	`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(book.VBID, book.Title, book.Author, book.UserID, book.UpdatedAt, book.ID)

	//result.RowsAffected()

	return err
}

func (book *Book) Delete() error {
	query := `
	DELETE FROM books
	WHERE id = ?
	`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(book.ID)

	//result.RowsAffected()

	return err
}
