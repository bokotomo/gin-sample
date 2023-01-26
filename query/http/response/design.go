package response

import (
	"gin-sample/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseDesignIndex is
func ResponseDesignIndex(context *gin.Context, designs *[10]domain.Design, total *int, err error) {
	if err != nil {
		ErrorResponse(context, http.StatusBadRequest, 0, err)
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

	context.JSON(http.StatusOK, gin.H{
		"designs": designsJson,
		"total":   &total,
	})
}

// ResponseDesignShow is
func ResponseDesignShow(context *gin.Context, design *domain.Design, err error) {
	if err != nil {
		ErrorResponse(context, http.StatusBadRequest, 0, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":        design.Id(),
		"title":     design.Title(),
		"text":      design.Text(),
		"good":      design.Good(),
		"comments":  design.Comments(),
		"date":      design.Date(),
		"thumbnail": design.Thumbnail(),
	})
}
