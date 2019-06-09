package response

import (
	// . "gin-sample/domain"
	"github.com/gin-gonic/gin"
)

func ResponseLoginIndex(context *gin.Context, token *string, err error) {
	if err != nil {
		ErrorResponse(context, 400, 0, err)
		return
	}

	context.JSON(200, gin.H{
		"token": &token,
	})
}
