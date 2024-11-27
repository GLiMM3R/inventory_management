package variant

import (
	"github.com/labstack/echo"
)

func InitVariantRoutes(e *echo.Echo, service VariantService) {
	h := NewProductHandler(service)
	r := e.Group("/products")
	r.POST("/:product_id/variants", h.AddVariant)
	r.PATCH("/:product_id/variants/:id", h.Update)
	r.DELETE("/:product_id/variants/:id", h.Delete)
}
