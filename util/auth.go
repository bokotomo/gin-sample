package util

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// VerifyBytes is
var VerifyBytes []byte

// SignKey is
var SignKey *rsa.PrivateKey

// FileInit is
func FileInit() {
	var signBytes []byte
	var err error
	VerifyBytes, err = ioutil.ReadFile("./infrastructure/ssh/public.sample")
	if err != nil {
		panic(err)
	}
	signBytes, err = ioutil.ReadFile("./infrastructure/ssh/secret.sample")
	if err != nil {
		panic(err)
	}
	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}
}

// CreateToken is
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

// CreatePasswordHash is
func CreatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// PasswordVerify パスワードがハッシュにマッチするかどうかを調べる
func PasswordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
