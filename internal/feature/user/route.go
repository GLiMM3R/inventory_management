package user

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitUserRoutes(e *echo.Echo, service UserServiceImpl) {
	h := NewUserHandler(service)
	r := e.Group("/users")
	r.Use(middleware.JWTAccessMiddleware)
	r.GET("", h.GetUsers)
	r.GET("/:username", h.GetUserByUsername)
	r.POST("", h.CreateUser)
}
