package inventory

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitInventoryRoutes(e *echo.Echo, service InventoryServiceImpl) {
	h := NewInventoryHandler(service)
	r := e.Group("/inventories")
	r.Use(middleware.JWTAccessMiddleware)
	r.GET("", h.GetInventories)
	r.GET("/:id", h.GetInventoryByID)
	r.POST("", h.CreateInventory)
	r.PATCH("/:id", h.UpdateInventory)
}
