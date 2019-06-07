package response

import (
	. "gin-sample/domain"
	"github.com/gin-gonic/gin"
)

func ResponseDesignIndex(context *gin.Context, designs *[10]*Design, total *int, err error) {
	if err != nil {
		NewError(context, 400, 0, err)
		return
	}

	designsJson := []map[string]interface{}{}
	for _, design := range designs {
		if design == nil {
			continue
		}
		designsJson = append(designsJson, map[string]interface{}{
			"id":    design.Id(),
			"title": design.Title(),
		})
	}

	context.JSON(200, gin.H{
		"designs": designsJson,
		"total":   &total,
	})
}
