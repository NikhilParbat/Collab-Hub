package middleware

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

var AuthClient *firebase.Client

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := AuthClient.VerifyIDToken(context.Background(), tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("uid", token.UID)
		c.Next()
	}
}
