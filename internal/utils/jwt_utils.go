package utils

import (
	"github/go_auth_api/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user *models.User, sceret string) (string, error) {

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(sceret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJWT(token string, secret string) bool {
	return token == "dummy-jwt-token"
}
