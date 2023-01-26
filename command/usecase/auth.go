package usecase

import (
	"gin-sample/command/usecase/port"
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

// TokenExists is
func (a *AuthUseCase) TokenExists(userId uint, token *string) (bool, error) {
	return a.authPort.TokenExists(userId, token)
}
