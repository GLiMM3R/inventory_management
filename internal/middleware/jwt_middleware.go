package middleware

import (
	"fmt"
	"inverntory_management/internal/exception"
	"inverntory_management/internal/service"
	"inverntory_management/internal/types"
	"strings"

	"github.com/labstack/echo"
)

func JWTAccessMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return exception.HandleError(c, exception.ErrUnauthorized)
		}

		// Assuming the Authorization header is in the format "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return exception.HandleError(c, exception.ErrInvalidToken)
		}

		claims, err := service.VerifyToken(tokenString, types.AccessToken)
		if err != nil {
			return exception.HandleError(c, err)
		}
		fmt.Println("claims=>", claims)
		c.Set("user", claims)

		return next(c)
	}
}

func JWTRefreshMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return exception.HandleError(c, exception.ErrUnauthorized)
		}

		// decodedToken, err := service.DecodeFromSHA256(authHeader)

		// Assuming the Authorization header is in the format "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return exception.HandleError(c, exception.ErrInvalidToken)
		}

		claims, err := service.VerifyToken(tokenString, types.RefreshToken)
		if err != nil {
			return exception.HandleError(c, err)
		}

		c.Set("user", claims)

		return next(c)
	}
}
