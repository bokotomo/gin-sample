package controller

import (
	"gin-sample/command/http/response"
	"gin-sample/command/repository"
	"gin-sample/command/usecase"

	"github.com/gin-gonic/gin"
)

// UserStore is
func UserStore(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	uc := usecase.NewUserUseCase(repository.NewUserRepository())
	token, err := uc.CreateUser(email, password)
	response.ResponseUserStore(c, token, err)
}
