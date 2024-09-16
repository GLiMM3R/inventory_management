package inventory_transfer

type InventoryTransferCreateDto struct {
	InventoryID string `json:"inventory_id"`
	ToBranchID  string `json:"to_branch_id"`
	Quantity    int    `json:"quantity"`
}

type InventoryTransferResponse struct {
	TransferID    string `json:"transfer_id"`
	InventoryID   string `json:"inventory_id"`
	InventoryName string `json:"inventory_name"`
	ToBranchID    string `json:"to_branch_id"`
	ToBranch      string `json:"to_branch"`
	Quantity      int    `json:"quantity"`
	TransferDate  int64  `json:"transfer_date"`
}
