package inventory_transfer

type InventoryTransferCreateDto struct {
	InventoryID  string `json:"inventory_id"`
	FromBranchID string `json:"from_branch_id"`
	ToBranchID   string `json:"to_branch_id"`
	Quantity     int    `json:"quantity"`
}
