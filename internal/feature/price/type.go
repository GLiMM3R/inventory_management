package price

type PriceCreateDto struct {
	InventoryID   string  `json:"inventory_id"`
	Price         float64 `json:"price"`
	EffectiveDate int64   `json:"effective_date"`
}

type PriceUpdateDto struct {
	Price         *float64 `json:"price"`
	EffectiveDate *int64   `json:"effective_date"`
}
