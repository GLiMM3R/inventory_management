package auth

import (
	"context"
	"fmt"
	"inverntory_management/config"
	"inverntory_management/internal/exception"
	"inverntory_management/internal/feature/user"
	"inverntory_management/internal/service"
	"inverntory_management/internal/types"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl interface {
	Login(request *AuthRequest) (*AuthResponse, error)
	Logout(token string) error
	GetRefreshToken(token string, userClaims *types.UserClaims) (*RefreshResponse, error)
}

type authService struct {
	userRepo    user.UserRepositoryImpl
	redisClient *redis.Client
}

func NewAuthService(userRepo user.UserRepositoryImpl, redisClient *redis.Client) AuthServiceImpl {
	return &authService{userRepo: userRepo, redisClient: redisClient}
}

// GetRefreshToken implements AuthServiceImpl.
func (s *authService) GetRefreshToken(token string, user *types.UserClaims) (*RefreshResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tokenString := strings.TrimPrefix(token, "Bearer ")
	if tokenString == token {
		return nil, exception.ErrInvalidToken
	}

	exists, err := s.redisClient.Exists(ctx, "refresh:"+tokenString).Result()
	if err != nil || exists == 0 {
		return nil, exception.ErrTokenExpired
	}

	accessToken, err := service.GenerateAccessToken(types.TokenPayload{UserID: user.Subject, Username: user.Username})
	if err != nil {
		fmt.Println("error here")
		return nil, exception.ErrInternal
	}

	// if want to rotate refresh token

	// refreshToken, err := service.GenerateRefreshToken(user.Subject)
	// if err != nil {
	// 	return nil, exception.ErrInternal
	// }

	// if err := s.redisClient.Del(ctx, "refresh:"+token).Err(); err != nil {
	// 	return nil, exception.ErrInternal
	// }

	// if err := s.redisClient.Set(ctx, "refresh:"+refreshToken, "active", time.Duration(config.AppConfig.REFRESH_EXPIRATION)*time.Second).Err(); err != nil {
	// 	return nil, exception.ErrInternal
	// }

	expiresIn := time.Now().Add(time.Duration(config.AppConfig.ACCESS_EXPIRATION) * time.Second).Unix()

	return &RefreshResponse{
		AccessToken: accessToken,
		ExpiresIn:   expiresIn,
	}, nil
}

// Login implements AuthServiceImpl.
func (s *authService) Login(request *AuthRequest) (*AuthResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user, err := s.userRepo.FindByUsername(request.Username)
	if err != nil {
		return nil, exception.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, exception.ErrInvalidCredentials
	}

	accessToken, err := service.GenerateAccessToken(types.TokenPayload{UserID: user.UserID, Username: user.Username})
	if err != nil {
		return nil, exception.ErrInternal
	}

	refreshToken, err := service.GenerateRefreshToken(user.UserID)
	if err != nil {
		return nil, exception.ErrInternal
	}

	if err := s.redisClient.Set(ctx, "refresh:"+refreshToken, "active", time.Duration(config.AppConfig.REFRESH_EXPIRATION)*time.Second).Err(); err != nil {
		return nil, exception.ErrInternal
	}

	expiresIn := time.Now().Add(time.Duration(config.AppConfig.ACCESS_EXPIRATION) * time.Second).Unix()

	return &AuthResponse{
		User: UserInfo{
			Username: user.Username,
			Email:    user.Email,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}, nil
}

// Logout implements AuthServiceImpl.
func (s *authService) Logout(token string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := s.redisClient.Del(ctx, "refresh:"+token).Err(); err != nil {
		return exception.ErrInternal
	}

	return nil
}
