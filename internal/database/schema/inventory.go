package schema

import (
	"gorm.io/gorm"
)

type Inventory struct {
	ID          uint           `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	InventoryID string         `json:"inventory_id" gorm:"primaryKey;unique;column:inventory_id"`
	BranchID    string         `json:"fk_branch_id" gorm:"column:fk_branch_id;uniqueIndex:idx_branch_id_name_sku"`
	Branch      Branch         `json:"-" gorm:"foreignKey:fk_branch_id;references:branch_id"`
	Name        string         `json:"name" gorm:"column:name;uniqueIndex:idx_branch_id_name_sku"`
	SKU         string         `json:"sku" gorm:"column:sku;uniqueIndex:idx_branch_id_name_sku"`
	Quantity    int            `json:"quantity" gorm:"column:quantity"`
	Price       float64        `json:"price" gorm:"column:price"`
	Status      string         `json:"status" gorm:"column:status"`
	CreatedAt   int64          `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   int64          `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
	Prices      []Price        `json:"-" gorm:"foreignKey:fk_inventory_id;references:inventory_id"`
}
