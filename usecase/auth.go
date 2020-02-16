package usecase

import (
	"gin-sample/usecase/port"
)

// AuthUseCase is
type AuthUseCase struct {
	authPort port.AuthPort
}

// NewAuthUseCase is
func NewAuthUseCase(authPort port.AuthPort) *AuthUseCase {
	return &AuthUseCase{authPort: authPort}
}

// Login is
func (a *AuthUseCase) Login(email, password string) (*string, error) {
	return a.authPort.Login(email, password)
}
