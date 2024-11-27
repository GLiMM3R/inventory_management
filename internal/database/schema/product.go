package schema

import (
	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	ProductID   string         `json:"product_id" gorm:"primaryKey;unique;column:product_id"`
	Name        string         `json:"name" gorm:"column:name"`
	CategoryID  string         `json:"category_id" gorm:"column:fk_category_id"`
	Category    Category       `json:"category" gorm:"foreignKey:fk_category_id;references:category_id"`
	Description string         `json:"description" gorm:"column:description;null"`
	CreatedAt   int64          `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   int64          `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
	Variants    []Variant      `json:"varaints" gorm:"foreignKey:fk_product_id;references:product_id"`
	Images      []ProductMedia `json:"images" gorm:"foreignKey:fk_product_id;references:product_id;"`
}

type ProductMedia struct {
	ID        uint           `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	MediaID   string         `json:"media_id" gorm:"unique;column:fk_media_id;index"`
	ProductID string         `json:"product_id" gorm:"column:fk_product_id;index"`
	CreatedAt int64          `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt int64          `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
	Media     Media          `json:"media" gorm:"foreignKey:fk_media_id;references:media_id"`
}
