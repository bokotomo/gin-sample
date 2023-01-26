package usecase

import (
	"gin-sample/query/usecase/port"
)

// AuthUseCase is
type AuthUseCase struct {
	authPort port.AuthPort
}

// NewAuthUseCase is
func NewAuthUseCase(authPort port.AuthPort) *AuthUseCase {
	return &AuthUseCase{authPort: authPort}
}

// TokenExists is
func (a *AuthUseCase) TokenExists(userId uint, token *string) (bool, error) {
	return a.authPort.TokenExists(userId, token)
}
