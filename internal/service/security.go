package service

import (
	"inverntory_management/config"
	"inverntory_management/internal/exception"
	"inverntory_management/internal/types"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(payload types.TokenPayload) (string, error) {
	secret := config.AppConfig.ACCESS_SECRET

	claims := &types.UserClaims{}
	claims.Subject = strconv.FormatUint(uint64(payload.UserID), 10)
	claims.Name = payload.Username
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Duration(config.AppConfig.ACCESS_EXPIRATION) * time.Second).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(userID string) (string, error) {
	secret := config.AppConfig.REFRESH_SECRET

	claims := &types.UserClaims{}
	claims.Subject = userID
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Duration(config.AppConfig.REDIS_DB) * time.Second).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateResetPasswordToken(userID string) (string, error) {
	secret := config.AppConfig.RESET_SECRET

	claims := &types.UserClaims{}
	claims.Subject = userID
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Duration(config.AppConfig.RESET_EXPIRATION) * time.Second).Unix()
	claims.Type = "reset_password"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string, secretKey []byte) (*types.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &types.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exception.ErrSigningMethodFailed
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, exception.ErrInvalidToken
	}

	claims, ok := token.Claims.(*types.UserClaims)

	if !ok {
		return nil, exception.ErrParseClaimed
	}

	return claims, nil
}

func VerifyToken(tokenString string, tokenType string) (*types.UserClaims, error) {
	switch tokenType {
	case types.AccessToken:
		secretKey := []byte(config.AppConfig.ACCESS_SECRET)
		return verifyToken(tokenString, secretKey)
	case types.RefreshToken:
		secretKey := []byte(config.AppConfig.REFRESH_SECRET)
		return verifyToken(tokenString, secretKey)
	case types.ResetPasswordToken:
		secretKey := []byte(config.AppConfig.RESET_SECRET)
		return verifyToken(tokenString, secretKey)
	default:
		return nil, exception.ErrInvalidToken
	}
}
