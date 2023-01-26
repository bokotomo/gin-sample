package middleware

import (
	"gin-sample/query/repository"
	"gin-sample/query/usecase"

	// 変える
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"crypto/rsa"
	"fmt"
	"gin-sample/util"
)

// AuthMiddleware is
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
		// todo
		var userId uint = 1
		uc := usecase.NewAuthUseCase(repository.NewAuthRepository())
		tokenExists, err := uc.TokenExists(userId, &token)
		if err != nil || !tokenExists {
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
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return lookupPublicKey(), nil
	})
	if err != nil || token.Valid {
		return err
	}

	return nil
}

func lookupPublicKey() *rsa.PublicKey {
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(util.VerifyBytes)
	if err != nil {
		panic(err)
	}
	return verifyKey
}
