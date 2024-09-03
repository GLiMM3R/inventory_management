package inventory

import "github.com/labstack/echo"

func InitInventoryRoutes(e *echo.Echo, service InventoryServiceImpl) {
	h := NewInventoryHandler(service)
	r := e.Group("/inventories")
	r.GET("", h.GetInventories)
	r.GET("/:id", h.GetInventoryByID)
	r.POST("", h.CreateInventory)
	r.PATCH("/:id", h.UpdateBranch)
}
