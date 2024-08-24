package routes

import (
	"example.com/go-blog/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", user.SignUp)
	server.POST("/login", user.Login)
}