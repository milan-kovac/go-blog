package main

import (
	"example.com/go-blog/db"
	"example.com/go-blog/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8099")
}