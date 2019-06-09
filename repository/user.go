package repository

import (
	// . "gi,ample/domain"
	"errors"
	. "gin-sample/driver"
	"gin-sample/model"
	. "gin-sample/util"
	"time"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return new(UserRepository)
}

/*
 * ユーザ検索
 */
func (this *UserRepository) FindUserByEmail(email string, user *model.User) error {
	return DB.Model(&model.User{}).Where("email = ?", email).First(user).Error
}

/*
 * メールがすでに存在するか
 */
func (this *UserRepository) EmailExists(email string) (bool, error) {
	var count int
	// TODO countでなくもっと良い方法にしたい
	if err := DB.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count != 0, nil
}

/*
 * ユーザ登録
 * email  メール
 * password パスワード
 */
func (this *UserRepository) CreateUser(email string, password string) (*string, error) {
	emailExists, err := NewUserRepository().EmailExists(email)
	if err != nil {
		return nil, err
	}
	if emailExists {
		return nil, errors.New("すでに登録されています。")
	}

	hash, err := CreatePasswordHash("password")
	if err != nil {
		return nil, err
	}

	user := model.User{Email: email, Password: hash}
	DB.Create(&user)
	token, err := CreateToken()
	if err != nil {
		return nil, err
	}

	if err := NewAuthRepository().TokenStore(user.ID, token, time.Now()); err != nil {
		return nil, err
	}

	return token, nil
}
