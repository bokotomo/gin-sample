package controller

import (
	. "gin-sample/domain"
	. "gin-sample/http/response"
	. "gin-sample/repository"
	. "gin-sample/usecase"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DesignIndex(c *gin.Context) {
	var (
		total   int
		designs [10]Design
	)
	uc := NewDesignUseCase(NewDesignRepository())
	page, _ := strconv.Atoi(c.Query("page"))
	err := uc.FindDesigns(&designs, &total, page)
	ResponseDesignIndex(c, &designs, &total, err)
}

func DesignShow(c *gin.Context) {
	var (
		design Design
	)
	designId, _ := strconv.Atoi(c.Param("designId"))
	uc := NewDesignUseCase(NewDesignRepository())
	err := uc.FindDesign(&design, designId)
	ResponseDesignShow(c, &design, err)
}
