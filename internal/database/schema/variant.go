package schema

import "gorm.io/gorm"

type Variant struct {
	ID           uint           `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	VariantID    string         `json:"variant_id" gorm:"primaryKey;unique;column:variant_id"`
	ProductID    string         `json:"product_id" gorm:"column:fk_product_id;"`
	SKU          string         `json:"sku" gorm:"unique;column:sku;"`
	ImageID      *string        `json:"media_id" gorm:"column:fk_media_id;uniqueIndex:idx_media_id_product_id"`
	Price        float64        `json:"price" gorm:"column:price;default:0"`
	Quantity     int            `json:"quantity" gorm:"column:quantity;default:0"`
	RestockLevel int            `json:"restock_level" gorm:"column:restock_level;default:0"`
	IsActive     bool           `json:"is_active" gorm:"column:is_active;default:true"`
	Status       string         `json:"status" gorm:"column:status"`
	CreatedAt    int64          `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    int64          `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
	Image        *Media         `json:"image" gorm:"foreignKey:fk_media_id;references:media_id"`
	Attributes   []Attribute    `json:"attributes" gorm:"foreignKey:fk_variant_id;references:variant_id"`
}

type Attribute struct {
	ID          uint   `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	AttributeID string `json:"attribute_id" gorm:"primaryKey;unique;column:attribute_id"`
	VariantID   string `json:"variant_id" gorm:"index;column:fk_variant_id;"`
	Attribute   string `json:"attribute" gorm:"column:attribute"`
	Value       string `json:"value" gorm:"column:value"`
}
