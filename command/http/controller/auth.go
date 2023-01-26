package controller

import (
	"gin-sample/command/http/response"
	"gin-sample/command/repository"
	"gin-sample/command/usecase"

	"github.com/gin-gonic/gin"
)

// LoginIndex is
func LoginIndex(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	uc := usecase.NewAuthUseCase(repository.NewAuthRepository())
	token, err := uc.Login(email, password)
	response.ResponseLoginIndex(c, token, err)
}

// AuthorizeIndex is
func AuthorizeIndex(c *gin.Context) {
	// email := c.PostForm("email")
	// password := c.PostForm("password")
	// token = email + password

	token := "aok" //Authorize()

	// Println(signBytes)
	// Println(verifyBytes)

	response.ResponseLoginIndex(c, &token, nil)
}
