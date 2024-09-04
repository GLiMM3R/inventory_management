package schema

type Sale struct {
	ID          uint      `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	SaleID      string    `json:"sale_id" gorm:"primaryKey;unique;column:sale_id"`
	InventoryID string    `json:"inventory_id" gorm:"column:fk_inventory_id;uniqueIndex:idx_inventory_id"`
	Inventory   Inventory `json:"-" gorm:"foreignKey:fk_inventory_id;references:inventory_id"`
	Quantity    int       `json:"quantity" gorm:"column:quantity"`
	TotalPrice  float64   `json:"total_price" gorm:"column:total_price"`
	SaleDate    int64     `json:"sale_date" gorm:"column:sale_date"`
	CreatedAt   int64     `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   int64     `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
}
