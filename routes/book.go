package routes

import (
	"bookstore-go/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getBooks(context *gin.Context) {
	books, err := models.AllBooks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch books."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Books fetched successfully.", "books": books})
}

func getBook(context *gin.Context) {
	bookID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data."})
		return
	}

	book, err := models.BookByID(bookID)
	if err != nil || book == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch book."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Book fetched successfully.", "book": book})
}

func createBook(context *gin.Context) {
	var book models.Book

	err := context.ShouldBindJSON(&book)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not bind the request data."})
		return
	}

	retrievedBook, err := models.BookByVBID(book.VBID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create book."})
		return
	}

	if retrievedBook != nil {
		context.JSON(http.StatusCreated, gin.H{"message": "The VBID exists."})
		return
	}

	userID, _ := context.Get("user_id")
	book.UserID = userID.(int64)

	err = book.Create()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create book."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Book created successfully.", "book": book})
}

func updateBook(context *gin.Context) {
	bookID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data."})
		return
	}

	retrievedBook, err := models.BookByID(bookID)
	if err != nil || retrievedBook == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch book."})
		return
	}

	var book models.Book
	err = context.ShouldBindJSON(&book)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not bind the request data."})
		return
	}

	retrievedBook, err = models.BookByVBID(book.VBID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create book."})
		return
	}

	if retrievedBook != nil && retrievedBook.ID != bookID {
		context.JSON(http.StatusCreated, gin.H{"message": "VBID is in use."})
		return
	}

	book.ID = bookID
	err = book.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update book."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Book updated successfully.", "book": book})
}

func deleteBook(context *gin.Context) {
	bookID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data."})
		return
	}

	retrievedBook, err := models.BookByID(bookID)
	if err != nil || retrievedBook == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch book."})
		return
	}

	err = retrievedBook.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete book."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully."})
}
