package schema

type Price struct {
	ID            uint      `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	PriceID       string    `json:"price_id" gorm:"primaryKey;unique;column:price_id"`
	InventoryID   string    `json:"inventory_id" gorm:"column:fk_inventory_id"`
	Inventory     Inventory `json:"-" gorm:"foreignKey:fk_inventory_id;references:inventory_id"`
	Price         float64   `json:"price" gorm:"column:price"`
	EffectiveDate int64     `json:"effective_date" gorm:"column:effective_date"`
	CreatedAt     int64     `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt     int64     `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
}
