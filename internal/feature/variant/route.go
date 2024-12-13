package variant

import (
	"github.com/labstack/echo"
)

func InitVariantRoutes(e *echo.Echo, service VariantService) {
	h := NewVariantHandler(service)
	r := e.Group("/variants")
	r.GET("/:id", h.GetVariant)
	r.POST("/:product_id", h.CreateVariant)
	r.PATCH("/:id", h.UpdateVariant)
	r.DELETE("/:id", h.DeleteVariant)
}
