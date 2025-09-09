package controllers

import (
	"github.com/MohdMusaiyab/blog-app/utils"
	"github.com/gin-gonic/gin"
)

// ProtectedTest is a dummy handler to test JWT middleware
func ProtectedTest(c *gin.Context) {
	// Extract user info from context (set by middleware)
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")

	utils.SuccessResponse(c, "Protected Route", gin.H{
		"user_id": userID,
		"email":   email,
	})
}
