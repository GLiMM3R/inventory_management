package user

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitUserRoutes(e *echo.Echo, service UserServiceImpl) {
	h := NewUserHandler(service)
	r := e.Group("/users")
	r.POST("", h.CreateUser)

	protected := e.Group("/restricted")
	protected.Use(middleware.JWTAccessMiddleware)
	protected.GET("", h.GetUsers)
	protected.GET("/:username", h.GetUserByUsername)
}
