package schema

type InventoryTransfer struct {
	ID           uint      `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	TransferID   string    `json:"transfer_id" gorm:"primaryKey;unique;column:transfer_id"`
	InventoryID  string    `json:"inventory_id" gorm:"column:fk_inventory_id"`
	Inventory    Inventory `json:"inventory" gorm:"foreignKey:fk_inventory_id;references:inventory_id"`
	FromBranchID string    `json:"from_branch_id" gorm:"column:fk_from_branch_id"`
	FromBranch   Branch    `json:"-" gorm:"foreignKey:fk_from_branch_id;references:branch_id"`
	ToBranchID   string    `json:"to_branch_id" gorm:"column:fk_to_branch_id"`
	ToBranch     Branch    `json:"-" gorm:"foreignKey:fk_to_branch_id;references:branch_id"`
	Quantity     int       `json:"quantity" gorm:"column:quantity"`
	TransferDate int64     `json:"transfer_date" gorm:"column:sale_date"`
	CreatedAt    int64     `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    int64     `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
}
