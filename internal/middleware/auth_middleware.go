package middleware

import (
	"net/http"
	"strings"

	"trinity/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks if the request contains a valid JWT token for user authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// Validate and parse the JWT token
		claims, err := utils.ParseJWTToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set the device id in the context for further use
		email := claims.Subject
		if email == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user"})
			c.Abort()
			return
		}

		c.Set("email", email)
		c.Next()
	}
}
