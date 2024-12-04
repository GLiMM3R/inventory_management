package schema

import (
	"gorm.io/gorm"
)

type Product struct {
	ID          uint             `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	ProductID   string           `json:"product_id" gorm:"primaryKey;unique;column:product_id"`
	Name        string           `json:"name" gorm:"column:name"`
	Description string           `json:"description" gorm:"column:description;null"`
	BasePrice   float64          `json:"base_price" gorm:"column:base_price;default:0"`
	CategoryID  string           `json:"category_id" gorm:"column:fk_category_id"`
	ThumbnailID *string          `json:"thumbnail_id" gorm:"column:thumbnail_id"`
	IsActive    bool             `json:"is_active" gorm:"column:is_active;default:true"`
	CreatedAt   int64            `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   int64            `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt   gorm.DeletedAt   `json:"deleted_at" gorm:"index;column:deleted_at"`
	Category    Category         `json:"category" gorm:"foreignKey:fk_category_id;references:category_id"`
	Variants    []ProductVariant `json:"variants" gorm:"foreignKey:product_id;references:product_id"`
	Thumbnail   *Media           `json:"thumbnail" gorm:"foreignKey:thumbnail_id;references:id"`
}

type ProductVariant struct {
	ID              uint           `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	VariantID       string         `json:"variant_id" gorm:"primaryKey;unique;column:variant_id"`
	ProductID       string         `json:"product_id" gorm:"column:product_id"`
	SKU             string         `json:"sku" gorm:"unique;column:sku"`
	VariantName     string         `json:"variant_name" gorm:"unique;column:variant_name"`
	AdditionalPrice float64        `json:"additional_price" gorm:"column:additional_price;default:0"`
	StockQuantity   int            `json:"stock_quantity" gorm:"column:stock_quantity;default:0"`
	RestockLevel    int            `json:"restock_level" gorm:"column:restock_level;default:0"`
	IsActive        bool           `json:"is_active" gorm:"column:is_active;default:true"`
	Status          string         `json:"status" gorm:"column:status"`
	CreatedAt       int64          `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt       int64          `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
	ImageID         *string        `json:"image_id" gorm:"column:image_id;uniqueIndex:idx_image_id"`
	Image           *Media         `json:"image" gorm:"foreignKey:image_id;references:id"`
	Attributes      []Attribute    `json:"attributes" gorm:"foreignKey:variant_id;references:variant_id"`
}

type Attribute struct {
	VariantID      string `json:"variant_id" gorm:"primaryKey;column:variant_id;foreignKey:variant_id;references:variant_id"`
	AttributeName  string `json:"attribute_name" gorm:"primaryKey;column:attribute_name"`
	AttributeValue string `json:"attribute_value" gorm:"column:attribute_value"`
}
