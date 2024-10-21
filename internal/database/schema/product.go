package schema

import (
	"gorm.io/gorm"
)

type Images []string

type Product struct {
	ID          uint           `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	ProductID   string         `json:"product_id" gorm:"primaryKey;unique;column:product_id"`
	Images      Images         `json:"images" gorm:"column:images;serializer:json"`
	Name        string         `json:"name" gorm:"column:name"`
	CategoryID  string         `json:"category_id" gorm:"column:fk_category_id"`
	Category    Category       `json:"-" gorm:"foreignKey:fk_category_id;references:category_id"`
	Description string         `json:"description" gorm:"column:description;null"`
	CreatedAt   int64          `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   int64          `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
	Variants    []Variant      `json:"varaints" gorm:"foreignKey:fk_product_id;references:product_id"`
}

type Variant struct {
	ID         uint           `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	VariantID  string         `json:"variant_id" gorm:"primaryKey;unique;column:variant_id"`
	ProductID  string         `json:"product_id" gorm:"column:fk_product_id;"`
	SKU        string         `json:"sku" gorm:"unique;column:sku;"`
	Image      string         `json:"image" gorm:"column:image"`
	Status     string         `json:"status" gorm:"column:status"`
	CreatedAt  int64          `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt  int64          `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
	Attributes []Attribute    `json:"attributes" gorm:"foreignKey:fk_variant_id;references:variant_id"`
	Price      []PriceHistory `json:"price" gorm:"foreignKey:fk_variant_id;references:variant_id"`
}

type Attribute struct {
	ID          uint   `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	AttributeID string `json:"attribute_id" gorm:"primaryKey;unique;column:attribute_id"`
	VariantID   string `json:"variant_id" gorm:"index;column:fk_variant_id;"`
	Attribute   string `json:"attribute" gorm:"column:attribute"`
	Value       string `json:"value" gorm:"column:value"`
}
