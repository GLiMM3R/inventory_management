package auth

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitAuthRoutes(e *echo.Echo, service AuthServiceImpl) {
	h := NewAuthHandler(service)
	r := e.Group("/auth")
	r.POST("/login", h.Login)

	protected := r.Group("")
	protected.Use(middleware.JWTRefreshMiddleware)
	protected.POST("/logout", h.Logout)
	protected.GET("/refresh-token", h.GetRefreshToken)
}
