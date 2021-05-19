package middleware

import (
	"net/http"
	"os"
	"password-share-server-golang/src/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func EnsureAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No JWT token found"})
			return
		}

		claims := &util.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Your token is invalid"})
				return
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad request", "err": err})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Your token is invalid"})
			return
		}

		claims.User.Password = "stop being snoopy"

		c.Set("user", claims.User)
		c.Set("token", token)

		c.Next()
	}
}
