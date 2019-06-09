package middleware

import (
	. "gin-sample/repository"
	. "gin-sample/usecase"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	// "github.com/dgrijalva/jwt-go/request"
	"crypto/rsa"
	. "fmt"
	. "gin-sample/util"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")
		if authorizationHeader == "" || authorizationHeader[:7] != "Bearer " {
			unauthorized(c)
			return
		}

		token := authorizationHeader[7:]
		if err := checkToken(&token); err != nil {
			unauthorized(c)
			return
		}

		Println("--------------------")
		uc := NewAuthUseCase(NewAuthRepository())
		token, err := uc.TokenExists(&token)
		if err != nil || token.Valid {
			unauthorized(c)
			return
		}
		c.Next()
	}
}

func unauthorized(context *gin.Context) {
	context.JSON(400, gin.H{
		"message": "認証失敗",
	})
}

func checkToken(tokenString *string) error {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isSame := token.Method.(*jwt.SigningMethodRSA); !isSame {
			return nil, Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return lookupPublicKey(), nil
	})
	if err != nil || token.Valid {
		return err
	}

	return nil
}

func lookupPublicKey() *rsa.PublicKey {
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(VerifyBytes)
	if err != nil {
		panic(err)
	}
	return verifyKey
}
