package controller

import (
	. "gin-sample/http/response"
	. "gin-sample/repository"
	. "gin-sample/usecase"
	"github.com/gin-gonic/gin"
)

func UserStore(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	uc := NewUserUseCase(NewUserRepository())
	token, err := uc.CreateUser(email, password)
	ResponseUserStore(c, token, err)
}
