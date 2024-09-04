package inventory_transfer

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitInventoryTransferRoutes(e *echo.Echo, service InventoryTransferServiceImpl) {
	h := NewInventoryTransactionHandler(service)
	r := e.Group("/inventory_transfers")
	r.Use(middleware.JWTAccessMiddleware)
	r.GET("", h.GetInventoryTransfers)
	r.GET("/:id", h.GetInventoryTransferByID)
	r.POST("", h.CreateInventoryTransfer)
}
