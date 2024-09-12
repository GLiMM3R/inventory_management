package auth

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitAuthRoutes(e *echo.Echo, service AuthServiceImpl) {
	h := NewAuthHandler(service)
	r := e.Group("/auth")
	r.POST("/login", h.Login)

	// protected := r.Group("")
	r.Use(middleware.JWTRefreshMiddleware)
	r.POST("/logout", h.Logout)
	r.GET("/refresh-token", h.GetRefreshToken)
}
