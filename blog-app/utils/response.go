package utils

import "github.com/gin-gonic/gin"

// Response structure for all APIs
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omit if nil
}

// SuccessResponse sends a successful response
func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(200, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error response with code
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, APIResponse{
		Success: false,
		Message: message,
		Data:    nil,
	})
}
