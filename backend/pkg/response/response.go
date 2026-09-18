package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, gin.H{"success": true, "message": msg, "data": data})
}
func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"success": false, "message": msg})
}

// SuccessResponse represents a success response format
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error response format
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
