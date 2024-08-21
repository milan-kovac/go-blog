package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var newUser NewUser
	err := context.ShouldBindJSON(&newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid data."})
		return
	}

	err = newUser.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "The email address is taken."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}