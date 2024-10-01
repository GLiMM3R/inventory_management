package product

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitProductRoutes(e *echo.Echo, service ProductServiceImpl) {
	h := NewProductHandler(service)
	r := e.Group("/products")
	r.Use(middleware.JWTAccessMiddleware)
	r.GET("", h.GetProducts)
	r.GET("/:id", h.GetProduct)
	r.POST("", h.CreateProduct)
}
