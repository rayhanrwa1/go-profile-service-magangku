package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"data":    nil,
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
				"data":    nil,
			})
			return
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"data":    nil,
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set("user_id", claims["user_id"])
		c.Set("userable_type", claims["userable_type"])

		c.Next()
	}
}
