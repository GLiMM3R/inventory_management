package inventory

type InventoryCreateDto struct {
	ProductID    string `json:"product_id" validate:"required"`
	Quantity     int    `json:"quantity"`
	RestockLevel int    `json:"restock_level"`
}

type InventoryUpdateDto struct {
	Quantity     *int  `json:"quantity,omitempty"`
	RestockLevel *int  `json:"restock_level,omitempty"`
	IsActive     *bool `json:"is_active,omitempty"`
}

var (
	ACTIVE     string = "active"
	DEPRECATED string = "deprecated"
	SOLD       string = "sold"
)
