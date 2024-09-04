package middleware

import (
	"inverntory_management/internal/service"
	"inverntory_management/internal/types"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func JWTAccessMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, map[string]string{"message": "token not found"})
		}

		// Assuming the Authorization header is in the format "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return echo.NewHTTPError(http.StatusUnauthorized, map[string]string{"message": "invalid token format"})
		}

		claims, err := service.VerifyToken(tokenString, types.AccessToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("user", claims)

		return next(c)
	}
}

func JWTRefreshMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, map[string]string{"message": "token not found"})
		}

		// Assuming the Authorization header is in the format "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return echo.NewHTTPError(http.StatusUnauthorized, map[string]string{"message": "invalid token format"})
		}

		claims, err := service.VerifyToken(tokenString, types.RefreshToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("user", claims)

		return next(c)
	}
}
