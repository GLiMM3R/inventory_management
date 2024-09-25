package schema

import "gorm.io/gorm"

type Inventory struct {
	ID           uint           `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	InventoryID  string         `json:"inventory_id" gorm:"primaryKey;unique;column:inventory_id"`
	BranchID     string         `json:"branch_id" gorm:"column:fk_branch_id;uniqueIndex:idx_branch_id_product_id"`
	Branch       Branch         `json:"-" gorm:"foreignKey:fk_branch_id;references:branch_id"`
	ProductID    string         `json:"product_id" gorm:"column:fk_product_id;uniqueIndex:idx_branch_id_product_id"`
	Product      Product        `json:"product" gorm:"foreignKey:fk_product_id;references:product_id"`
	Quantity     int            `json:"quantity" gorm:"column:quantity"`
	RestockLevel int            `json:"restock_level" gorm:"column:restock_level"`
	IsActive     bool           `json:"is_active" gorm:"column:is_active"`
	CreatedAt    int64          `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    int64          `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
}
