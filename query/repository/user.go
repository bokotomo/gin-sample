package repository

import (
	"gin-sample/driver"
	"gin-sample/model"
)

// UserRepository is
type UserRepository struct{}

// NewUserRepository is
func NewUserRepository() *UserRepository {
	return new(UserRepository)
}

// FindUserByEmail ユーザ検索
func (u *UserRepository) FindUserByEmail(email string, user *model.User) error {
	return driver.DB.Model(&model.User{}).Where("email = ?", email).First(user).Error
}
