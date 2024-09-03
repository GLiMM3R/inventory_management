package user

import "github.com/labstack/echo"

func InitUserRoutes(e *echo.Echo, service UserServiceImpl) {
	h := NewUserHandler(service)
	r := e.Group("/users")
	r.GET("", h.GetUsers)
	r.GET("/:username", h.GetUserByUsername)
	r.POST("", h.CreateUser)
}
