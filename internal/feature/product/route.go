package product

import (
	"github.com/labstack/echo"
)

func InitProductRoutes(e *echo.Echo, service ProductServiceImpl) {
	h := NewProductHandler(service)
	r := e.Group("/products")
	r.GET("", h.GetProducts)
	// r.Use(middleware.JWTAccessMiddleware)
	r.GET("/:id", h.GetProduct)
	r.POST("", h.CreateProduct)
	r.POST("/:id/variants", h.AddVariant)
}
