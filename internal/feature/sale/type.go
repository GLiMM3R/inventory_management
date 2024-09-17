package sale

type SaleCreateDto struct {
	Items []Item `json:"items"`
}

type Item struct {
	InventoryID string `json:"inventory_id"`
	Quantity    int    `json:"quantity"`
}
