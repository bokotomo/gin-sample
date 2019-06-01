package controller

import (
  "github.com/gin-gonic/gin"
  . "gin-sample/http/response"
  . "gin-sample/usecase"
  . "gin-sample/repository"
)

func DesignIndex(context *gin.Context) {
	uc := NewDesignUseCase(NewDesignRepository())
	designs, err := uc.FindAllDesigns()
  ResponseDesignIndex(context, designs, err)
}
