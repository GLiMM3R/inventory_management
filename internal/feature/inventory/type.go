package inventory

type InventoryCreateDto struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type InventoryUpdateDto struct {
	Name     *string  `json:"name,omitempty"`
	Quantity *int     `json:"quantity,omitempty"`
	Status   *string  `json:"status,omitempty"`
	Price    *float64 `json:"price,omitempty"`
}

var (
	ACTIVE     string = "active"
	DEPRECATED string = "deprecated"
	SOLD       string = "sold"
)
