package usecase

import (
	"gin-sample/usecase/port"
)

// UserUseCase is
type UserUseCase struct {
	userPort port.UserPort
}

// NewUserUseCase is
func NewUserUseCase(userPort port.UserPort) *UserUseCase {
	return &UserUseCase{userPort: userPort}
}

// CreateUser is
func (u *UserUseCase) CreateUser(email, password string) (*string, error) {
	return u.userPort.CreateUser(email, password)
}
