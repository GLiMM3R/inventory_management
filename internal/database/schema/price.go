package schema

type PriceHistory struct {
	ID            uint    `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	PriceID       string  `json:"price_id" gorm:"primaryKey;unique;column:price_id"`
	VariantID     string  `json:"variant_id" gorm:"column:fk_variant_id"`
	Variant       Variant `json:"-" gorm:"foreignKey:fk_variant_id;references:variant_id"`
	NewPrice      float64 `json:"new_price" gorm:"column:new_price"`
	OldPrice      float64 `json:"old_price" gorm:"column:old_price"`
	EffectiveDate int64   `json:"effective_date" gorm:"column:effective_date"`
	CreatedAt     int64   `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt     int64   `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
}
