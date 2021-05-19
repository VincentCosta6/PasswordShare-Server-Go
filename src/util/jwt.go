package util

import (
	"os"
	"password-share-server-golang/src/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	User *models.User
	jwt.StandardClaims
}

func CreateJWTTokenString(user *models.User) (string, error) {
	claims := &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Hour * 24).Unix(),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}
