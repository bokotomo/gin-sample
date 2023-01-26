package controller

import (
	"gin-sample/domain"
	"gin-sample/query/http/response"
	"gin-sample/query/repository"
	"gin-sample/query/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DesignIndex デザイン一覧を取得する
func DesignIndex(c *gin.Context) {
	var (
		total   int
		designs [10]domain.Design
	)

	uc := usecase.NewDesignUseCase(repository.NewDesignRepository())
	page, _ := strconv.Atoi(c.Query("page"))
	err := uc.FindDesigns(&designs, &total, page)

	response.ResponseDesignIndex(c, &designs, &total, err)
}

// DesignShow designIdのデザインを取得する
func DesignShow(c *gin.Context) {
	var design domain.Design
	designID, _ := strconv.Atoi(c.Param("designId"))

	uc := usecase.NewDesignUseCase(repository.NewDesignRepository())
	err := uc.FindDesign(&design, designID)

	response.ResponseDesignShow(c, &design, err)
}
