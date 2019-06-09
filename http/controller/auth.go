package controller

import (
	. "gin-sample/http/response"
	. "gin-sample/repository"
	. "gin-sample/usecase"
	"github.com/gin-gonic/gin"
)

func LoginIndex(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	uc := NewAuthUseCase(NewAuthRepository())
	token, err := uc.Login(email, password)
	ResponseLoginIndex(c, token, err)
}

func AuthorizeIndex(c *gin.Context) {
	// email := c.PostForm("email")
	// password := c.PostForm("password")
	// token = email + password

	token := "aok" //Authorize()

	// Println(signBytes)
	// Println(verifyBytes)

	ResponseLoginIndex(c, &token, nil)
}
