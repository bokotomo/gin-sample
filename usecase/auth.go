package usecase

import (
	. "gin-sample/usecase/port"
)

type AuthUseCase struct {
	authPort AuthPort
}

func NewAuthUseCase(authPort AuthPort) *AuthUseCase {
	return &AuthUseCase{authPort: authPort}
}

func (this *AuthUseCase) Login(email string, password string) (*string, error) {
	return this.authPort.Login(email, password)
}
