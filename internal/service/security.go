package service

import (
	"fmt"
	"inverntory_management/config"
	"inverntory_management/internal/exception"
	"inverntory_management/internal/types"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(payload types.TokenPayload) (string, error) {
	privateBytes := config.AppConfig.PRIVATE_KEY

	secret, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}

	claims := &types.UserClaims{}
	claims.Subject = payload.UserID
	claims.Username = payload.Username
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Duration(config.AppConfig.ACCESS_EXPIRATION) * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)

		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(userID string) (string, error) {
	secret := config.AppConfig.REFRESH_SECRET
	claims := &types.UserClaims{}
	claims.Subject = userID
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Duration(config.AppConfig.REFRESH_EXPIRATION) * time.Minute).Unix()

	fmt.Println(claims)

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
	claims.ExpiresAt = time.Now().Add(time.Duration(config.AppConfig.RESET_EXPIRATION) * time.Minute).Unix()
	claims.Type = "reset_password"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyAccessToken(tokenString string, secret []byte) (*types.UserClaims, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(secret)
	if err != nil {
		log.Fatalf("Error parsing public key: %v", err)
	}

	token, err := jwt.ParseWithClaims(tokenString, &types.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, exception.ErrSigningMethodFailed
		}

		return publicKey, nil
	})

	if err != nil {
		return nil, exception.ErrInvalidToken
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

func verifyToken(tokenString string, secretKey []byte) (*types.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &types.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exception.ErrSigningMethodFailed
		}
		// Ensure the signing method is HS256
		if token.Method != jwt.SigningMethodHS256 {
			return nil, exception.ErrSigningMethodFailed
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, exception.ErrInvalidToken
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
		secretKey := config.AppConfig.PUBLIC_KEY
		return verifyAccessToken(tokenString, secretKey)
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
