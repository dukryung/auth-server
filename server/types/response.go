package types

import "github.com/gin-gonic/gin"

func SetResponse(message string, status int, error bool, data interface{}) gin.H {
	return gin.H{
		"message": message,
		"status":  status,
		"error":   error,
		"data":    data,
	}
}
