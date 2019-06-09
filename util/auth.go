package util

import (
  jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	. "io/ioutil"
	"time"
  "crypto/rsa"
)

var VerifyBytes []byte
var SignKey *rsa.PrivateKey

func FileInit() {
  var signBytes []byte
  var err error
	VerifyBytes, err = ReadFile("./infrastructure/ssh/public.sample")
	if err != nil {
		panic(err)
	}
	signBytes, err = ReadFile("./infrastructure/ssh/secret.sample")
	if err != nil {
		panic(err)
	}
	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}
}

func CreateToken() (*string, error) {
	tokenRSA := jwt.New(jwt.SigningMethodRS256)
	claims := tokenRSA.Claims.(jwt.MapClaims)
	claims["name"] = "test"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, err := tokenRSA.SignedString(SignKey)
  if err != nil {
  	return nil, err
  }

	return &token, nil
}

func CreatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// パスワードがハッシュにマッチするかどうかを調べる
func PasswordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
