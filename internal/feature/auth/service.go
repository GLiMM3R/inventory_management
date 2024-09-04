package auth

import (
	"inverntory_management/internal/feature/user"
	"inverntory_management/internal/types"
)

type AuthServiceImpl interface {
	Login(username, password string) (*AuthResponse, error)
	Logout(token string) error
	GetRefreshToken(token string, userClaims *types.UserClaims) (*AuthResponse, error)
}

type authService struct {
	userRepo user.UserRepositoryImpl
}

// GetRefreshToken implements AuthServiceImpl.
func (s *authService) GetRefreshToken(token string, userClaims *types.UserClaims) (*AuthResponse, error) {
	panic("unimplemented")
}

// Login implements AuthServiceImpl.
func (s *authService) Login(username string, password string) (*AuthResponse, error) {
	panic("unimplemented")
}

// Logout implements AuthServiceImpl.
func (s *authService) Logout(token string) error {
	panic("unimplemented")
}

func NewAuthService(userRepo user.UserRepositoryImpl) AuthServiceImpl {
	return &authService{userRepo: userRepo}
}
