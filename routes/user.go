package routes

import (
	"bookstore-go/models"
	"bookstore-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {
	var loginUser models.LoginUser
	err := context.ShouldBindJSON(&loginUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data."})
		return
	}

	user, err := loginUser.Authenticate()

	// if reset pw is required, return reset pw
	if loginUser.MustChangePassword {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are required to change your password."})
		return
	}

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password."})
		return
	}

	token, err := utils.GenerateJWT(user.RoleID, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "token could not be generated"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successful.", "token": token})
}

func changePassword(context *gin.Context) {
	var loginUser models.LoginUser

	err := context.ShouldBindJSON(&loginUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// This should now print the new password
	if loginUser.NewPassword == "" || loginUser.NewPassword == loginUser.Password {
		context.JSON(http.StatusConflict, gin.H{"message": "New password should not be empty and should not be the same as old password."})
		return
	}

	user, err := loginUser.Authenticate()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	err = user.UpdatePassword(&loginUser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Changing of password is unsuccessful."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Password has been changed successfully."})
}

func register(context *gin.Context) {
	var registerUser models.RegisterUser

	err := context.ShouldBindJSON(&registerUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data."})
		return
	}

	registerUser.RoleID = 2 // manager role

	err = registerUser.Create()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user account."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}
