package response

import (
	// . "gin-sample/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseLoginIndex is
func ResponseLoginIndex(context *gin.Context, token *string, err error) {
	if err != nil {
		ErrorResponse(context, http.StatusBadRequest, 0, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token": &token,
	})
}
