package sale

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitSaleRoutes(e *echo.Echo, service SaleServiceImpl) {
	h := NewSaleHandler(service)
	r := e.Group("/sales")
	r.Use(middleware.JWTAccessMiddleware)
	r.GET("", h.GetSales)
	r.GET("/:id", h.GetPriceByID)
	r.POST("", h.CreateSale)
}
