package price

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitPriceRoutes(e *echo.Echo, service PriceServiceImpl) {
	h := NewPriceHandler(service)
	r := e.Group("/prices")
	r.Use(middleware.JWTAccessMiddleware)
	r.GET("", h.GetPrices)
	r.GET("/:id", h.GetPriceByID)
	r.POST("", h.CreatePrice)
	r.PATCH("/:id", h.UpdatePrice)
}
