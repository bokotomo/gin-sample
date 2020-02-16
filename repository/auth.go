package repository

import (
	. "gin-sample/driver"
	"gin-sample/model"
	. "gin-sample/util"
	"time"
)

// AuthRepository is
type AuthRepository struct{}

// NewAuthRepository is
func NewAuthRepository() *AuthRepository {
	return new(AuthRepository)
}

// Login ログイン
func (a *AuthRepository) Login(email string, password string) (*string, error) {
	user := model.User{}
	if err := NewUserRepository().FindUserByEmail(email, &user); err != nil {
		return nil, err
	}
	if err := PasswordVerify(user.Password, password); err != nil {
		return nil, err
	}

	token, err := CreateToken()
	if err != nil {
		return nil, err
	}

	if err := NewAuthRepository().TokenStore(user.ID, token, time.Now()); err != nil {
		return nil, err
	}

	return token, nil
}

// TokenStore token追加
func (a *AuthRepository) TokenStore(userId uint, token *string, exp time.Time) error {
	DB.Create(&model.Token{UserID: userId, Token: *token, Exp: exp})
	return nil
}

// TokenExists tokenがあるか
func (a *AuthRepository) TokenExists(userId uint, token *string) (bool, error) {
	var count int
	// TODO countでなくもっと良い方法にしたい
	if err := DB.Model(&model.Token{}).Where("user_id = ? AND token = ?", userId, token).Count(&count).Error; err != nil {
		return false, err
	}
	return count != 0, nil
}
