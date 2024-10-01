package product

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitProductRoutes(e *echo.Echo, service ProductServiceImpl) {
	h := NewProductHandler(service)
	r := e.Group("/prices")
	r.Use(middleware.JWTAccessMiddleware)
	r.POST("", h.CreateProduct)
}
