package repository

import (
	"errors"
	. "gin-sample/driver"
	"gin-sample/model"
	"gin-sample/util"
	"time"
)

// UserRepository is
type UserRepository struct{}

// NewUserRepository is
func NewUserRepository() *UserRepository {
	return new(UserRepository)
}

// FindUserByEmail ユーザ検索
func (u *UserRepository) FindUserByEmail(email string, user *model.User) error {
	return DB.Model(&model.User{}).Where("email = ?", email).First(user).Error
}

// EmailExists メールがすでに存在するか
func (u *UserRepository) EmailExists(email string) (bool, error) {
	var count int
	// TODO countでなくもっと良い方法にしたい
	if err := DB.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count != 0, nil
}

// CreateUser ユーザ登録
func (u *UserRepository) CreateUser(email string, password string) (*string, error) {
	emailExists, err := NewUserRepository().EmailExists(email)
	if err != nil {
		return nil, err
	}
	if emailExists {
		return nil, errors.New("すでに登録されています。")
	}

	hash, err := util.CreatePasswordHash("password")
	if err != nil {
		return nil, err
	}

	user := model.User{Email: email, Password: hash}
	DB.Create(&user)
	token, err := util.CreateToken()
	if err != nil {
		return nil, err
	}

	if err := NewAuthRepository().TokenStore(user.ID, token, time.Now()); err != nil {
		return nil, err
	}

	return token, nil
}
