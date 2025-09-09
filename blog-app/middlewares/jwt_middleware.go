package middlewares

import (
	"net/http"
	"strings"

	"github.com/MohdMusaiyab/blog-app/utils"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware validates JWT and attaches user info to context
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "authorization header missing")
			c.Abort()
			return
		}

		// Check format: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "invalid authorization format")
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parse token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "invalid or expired token")
			c.Abort()
			return
		}

		// Attach user info to context for downstream handlers
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)

		// Continue to next handler
		c.Next()
	}
}
