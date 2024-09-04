package sale

import "github.com/labstack/echo"

func InitSaleRoutes(e *echo.Echo, service SaleServiceImpl) {
	h := NewSaleHandler(service)
	r := e.Group("/prices")
	r.GET("", h.GetSales)
	r.GET("/:id", h.GetPriceByID)
	r.POST("", h.CreateSale)
}
