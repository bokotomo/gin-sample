package response

import (
	"github.com/gin-gonic/gin"
)

func ErrorResponse(context *gin.Context, statusCode int, errorCode int, err error) {
	context.JSON(statusCode, gin.H{
		"code":  errorCode,
		"error": err.Error(),
	})
}
