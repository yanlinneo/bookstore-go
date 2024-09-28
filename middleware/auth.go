package middleware

import (
	"bookstore-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorize(context *gin.Context) {
	token := context.GetHeader("Authorization")
	userID, roleID, err := utils.ValidateJWT(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized. Token might be invalid or have expired."})
		return
	}

	if context.Request.Method == http.MethodDelete {
		if roleID != 1 {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "You do not have the required permission to perform this action"})
			return
		}
	}

	context.Set("user_id", userID)
	context.Next()
}
