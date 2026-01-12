package middleware

import (
	"net/http"
	"strings"
	"uam-golang/internal/config"
	"uam-golang/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Remove Bearer prefix
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		config, _ := config.LoadConfig() // In real app, consider injecting or caching
		claims, err := utils.ValidateToken(tokenString, config.JWTSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
