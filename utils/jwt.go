package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey string = "learn90-onTHE_90"

func GenerateJWT(roleID int64, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role_id": roleID,
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func ValidateJWT(tokenString string) (int64, int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, 0, err
	}

	if !parsedToken.Valid {
		return 0, 0, fmt.Errorf("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, 0, fmt.Errorf("invalid token claims")
	}

	userID := int64(claims["user_id"].(float64))
	roleID := int64(claims["role_id"].(float64))

	return userID, roleID, nil
}
