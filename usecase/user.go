package usecase

import (
	. "gin-sample/usecase/port"
)

type UserUseCase struct {
	userPort UserPort
}

func NewUserUseCase(userPort UserPort) *UserUseCase {
	return &UserUseCase{userPort: userPort}
}

func (this *UserUseCase) CreateUser(email string, password string) (*string, error) {
	return this.userPort.CreateUser(email, password)
}
