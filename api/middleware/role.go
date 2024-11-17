package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JWT_SECRET = []byte("secret-key")

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Tokenni ajratib olish
		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			if err, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return err, nil
			}
			return JWT_SECRET, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not read claims"})
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"message": "You do not have the required role"})
			c.Abort()
			return
		}

		c.Next()
	}
}
