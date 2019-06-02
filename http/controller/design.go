package controller

import (
	. "gin-sample/http/response"
	. "gin-sample/repository"
	. "gin-sample/usecase"
	"github.com/gin-gonic/gin"
)

func DesignIndex(context *gin.Context) {
	uc := NewDesignUseCase(NewDesignRepository())
	designs, err := uc.FindAllDesigns()
	ResponseDesignIndex(context, designs, err)
}
