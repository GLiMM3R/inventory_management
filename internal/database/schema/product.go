package schema

import "gorm.io/gorm"

type Product struct {
	ID          uint           `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	ProductID   string         `json:"variant_id" gorm:"primaryKey;unique;column:variant_id"`
	Name        string         `json:"name" gorm:"column:name"`
	CategoryID  string         `json:"category_id" gorm:"column:fk_category_id"`
	Category    Category       `json:"-" gorm:"foreignKey:fk_category_id;references:category_id"`
	Description string         `json:"description" gorm:"column:description;null"`
	CreatedAt   int64          `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   int64          `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
}

type ProductVariant struct {
	ID        uint           `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	VariantID string         `json:"variant_id" gorm:"primaryKey;unique;column:variant_id"`
	ProductID string         `json:"product_id" gorm:"column:fk_product_id;"`
	Product   Product        `json:"-" gorm:"foreignKey:fk_product_id;references:product_id"`
	SKU       string         `json:"sku" gorm:"column:sku;unique;"`
	Status    string         `json:"status" gorm:"column:status"`
	CreatedAt int64          `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt int64          `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
}

type VariantOption struct {
	ID          uint   `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	OptionID    string `json:"option_id" gorm:"primaryKey;unique;column:option_id"`
	OptionName  string `json:"option_name" gorm:"column:option_name"`
	OptionValue string `json:"option_value" gorm:"column:option_value"`
}

type VariantOptionAssignment struct {
	ID        uint           `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	VariantID string         `json:"variant_id" gorm:"column:fk_variant_id;"`
	Variant   ProductVariant `json:"variant" gorm:"foreignKey:fk_variant_id;references:variant_id"`
	OptionID  string         `json:"option_id" gorm:"column:fk_option_id;"`
	Option    VariantOption  `json:"option" gorm:"foreignKey:fk_option_id;references:option_id"`
}
