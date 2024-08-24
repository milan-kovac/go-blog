package user

import (
	"net/http"

	"example.com/go-blog/shared"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var newUser UserData
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


func Login(context *gin.Context)  {
	var user UserData

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := shared.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})

	
}
