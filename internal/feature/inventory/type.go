package inventory

type InventoryCreateDto struct {
	BranchID string `json:"branch_id"`
	Name     string `json:"name"`
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

type InventoryUpdateDto struct {
	Name     *string `json:"name,omitempty"`
	SKU      *string `json:"sku,omitempty"`
	Quantity *int    `json:"quantity,omitempty"`
	Status   *string `json:"status,omitempty"`
}

var (
	ACTIVE     string = "active"
	DEPRECATED string = "deprecated"
	SOLD       string = "sold"
)
