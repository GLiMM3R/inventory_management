package sale

type SaleCreateDto struct {
	InventoryID string `json:"inventory_id"`
	Quantity    int    `json:"quantity"`
}
