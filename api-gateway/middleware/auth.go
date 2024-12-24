package middleware

import (
	pbAuth "20241224/auth-service/proto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(userClient pbAuth.AuthServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		req := &pbAuth.ValidateTokenRequest{Token: token}
		res, err := userClient.ValidateToken(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("user_id", res.UserId)
		c.Next()
	}
}
