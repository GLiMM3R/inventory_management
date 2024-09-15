package report

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitReportRoutes(e *echo.Echo, service ReportServiceImpl) {
	h := NewSaleHandler(service)
	r := e.Group("/reports")
	r.Use(middleware.JWTAccessMiddleware)
	r.GET("/sales", h.GetSalesReport)
}
