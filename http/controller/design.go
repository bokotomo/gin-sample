package controller

import (
	. "gin-sample/http/response"
	. "gin-sample/repository"
	. "gin-sample/usecase"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DesignIndex(c *gin.Context) {
	uc := NewDesignUseCase(NewDesignRepository())
	page, _ := strconv.Atoi(c.Query("page"))
	designs, err := uc.FindDesigns(page)
	ResponseDesignIndex(c, designs, err)
}
