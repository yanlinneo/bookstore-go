package routes

import (
	"bookstore-go/middleware"

	"github.com/gin-gonic/gin"
)

func Register(server *gin.Engine) {
	server.POST("/login", login)
	server.POST("/changePassword", changePassword)

	authenticate := server.Group("/")
	authenticate.Use(middleware.Authorize)

	authenticate.GET("/books", getBooks)
	authenticate.GET("/books/:id", getBook)
	authenticate.POST("/books", createBook)
	authenticate.PUT("/books/:id", updateBook)
	authenticate.DELETE("/books/:id", deleteBook) // only admin can perform this action

	authenticate.POST("/register", register)

}
