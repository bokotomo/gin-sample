package response

import (
	. "gin-sample/domain"
	"github.com/gin-gonic/gin"
)

func ResponseDesignIndex(context *gin.Context, designs []*Design, err error) {
	if err != nil {
		NewError(context, 400, 0, err)
		return
	}

	context.JSON(200, gin.H{
		"designs": designs,
	})
}
