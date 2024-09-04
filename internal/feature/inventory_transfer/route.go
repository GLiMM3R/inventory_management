package inventory_transfer

import "github.com/labstack/echo"

func InitInventoryTransferRoutes(e *echo.Echo, service InventoryTransferServiceImpl) {
	h := NewInventoryTransactionHandler(service)
	r := e.Group("/inventory_transfers")
	r.GET("", h.GetInventoryTransfers)
	r.GET("/:id", h.GetInventoryTransferByID)
	r.POST("", h.CreateInventoryTransfer)
}
