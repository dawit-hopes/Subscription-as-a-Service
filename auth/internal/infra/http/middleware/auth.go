package middleware

import (
	"net/http"
	"strings"

	"github.com/dawit_hopes/saas/auth/internal/domain/port/outbound"
	"github.com/gin-gonic/gin"
)

func TokenMiddleware(tokenProvider outbound.TokenProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		// Expect header format: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}

		token := parts[1]

		claims, err := tokenProvider.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Message})
			return
		}

		c.Set("userID", claims.Subject)

		c.Next()
	}
}
