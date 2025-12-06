package lib

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("rP9aL8sB#yT1gHj!WzM0nKdXe@u")

type CustomClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(id string, role string, username string) (string, error) {
	claims := CustomClaims{
		Id:       id,
		Role:     role,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "KageNoEn",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
