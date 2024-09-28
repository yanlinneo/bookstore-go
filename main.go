package main

import (
	"bookstore-go/database"
	"bookstore-go/models"
	"bookstore-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	models.CreateAdmin()
	server := gin.Default()

	routes.Register(server)
	server.Run(":8080")
}
