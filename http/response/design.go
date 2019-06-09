package response

import (
	. "gin-sample/domain"
	"github.com/gin-gonic/gin"
)

func ResponseDesignIndex(context *gin.Context, designs *[10]Design, total *int, err error) {
	if err != nil {
		ErrorResponse(context, 400, 0, err)
		return
	}

	designsJson := []map[string]interface{}{}
	for _, design := range designs {
		if design.NotExists() {
			continue
		}
		designsJson = append(designsJson, map[string]interface{}{
			"id":    design.Id(),
			"title": design.Title(),
			"good":  design.Good(),
		})
	}

	context.JSON(200, gin.H{
		"designs": designsJson,
		"total":   &total,
	})
}

func ResponseDesignShow(context *gin.Context, design *Design, err error) {
	if err != nil {
		ErrorResponse(context, 400, 0, err)
		return
	}

	context.JSON(200, gin.H{
		"id":        design.Id(),
		"title":     design.Title(),
		"text":      design.Text(),
		"good":      design.Good(),
		"comments":  design.Comments(),
		"date":      design.Date(),
		"thumbnail": design.Thumbnail(),
	})
}
