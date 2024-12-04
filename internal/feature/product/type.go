package product

import (
	"inverntory_management/internal/feature/media"
)

type CreateProductDTO struct {
	Name        string                `json:"name" validate:"required"`
	BasePrice   float64               `json:"base_price" validate:"required,min=0"`
	CategoryID  string                `json:"category_id" validate:"required"`
	Description string                `json:"description"`
	Thumbnail   *media.CreateMediaDTO `json:"thumbnail,omitempty"`
	Variants    []CreateVariantDTO    `json:"variants"`
}

type CreateVariantDTO struct {
	SKU             string                `json:"sku" validate:"required"`
	VariantName     string                `json:"variant_name" validate:"required"`
	AdditionalPrice float64               `json:"additional_price" validate:"min=0"`
	StockQuantity   int                   `json:"stock_quantity" validate:"min=0"`
	RestockLevel    int                   `json:"restock_level" validate:"min=0"`
	Image           *media.CreateMediaDTO `json:"image,omitempty"`
	Attributes      []CreateAttributeDTO  `json:"attributes"`
}

type CreateAttributeDTO struct {
	AttributeName  string `json:"attribute_name" validate:"required"`
	AttributeValue string `json:"attribute_value" validate:"required"`
}

type UpdateProductDTO struct {
	Name        *string               `json:"name" validate:"required"`
	BasePrice   *float64              `json:"base_price" validate:"required,min=0"`
	CategoryID  *string               `json:"category_id" validate:"required"`
	Description *string               `json:"description"`
	Thumbnail   *media.CreateMediaDTO `json:"thumbnail,omitempty"`
	IsActive    *bool                 `json:"is_active,omitempty"`
}

type ProductListResponse struct {
	ProductID    string  `json:"product_id"`
	Name         string  `json:"name"`
	BasePrice    float64 `json:"base_price"`
	ThumbnailURL string  `json:"thumbnail_url"`
	CategoryName string  `json:"category_name"`
	Description  string  `json:"description"`
	IsActive     bool    `json:"is_active"`
	CreatedAt    int64   `json:"created_at"`
	UpdatedAt    int64   `json:"updated_at"`
}

type VariantResponse struct {
	VariantID       string              `json:"variant_id"`
	SKU             string              `json:"sku"`
	VariantName     string              `json:"variant_name"`
	AdditionalPrice float64             `json:"additional_price"`
	StockQuantity   int                 `json:"stock_quantity"`
	RestockLevel    int                 `json:"restock_level"`
	IsActive        bool                `json:"is_active"`
	Status          string              `json:"status"`
	CreatedAt       int64               `json:"created_at"`
	UpdatedAt       int64               `json:"updated_at"`
	ImageURL        string              `json:"image_url"`
	Attributes      []AttributeResponse `json:"attributes"`
}

type AttributeResponse struct {
	AttributeName  string `json:"attribute_name"`
	AttributeValue string `json:"attribute_value"`
}

type ProductResponse struct {
	ProductID    string            `json:"product_id"`
	Name         string            `json:"name"`
	BasePrice    float64           `json:"base_price"`
	ThumbnailURL string            `json:"thumbnail_url"`
	CategoryID   string            `json:"category_id"`
	CategoryName string            `json:"category_name"`
	Description  string            `json:"description"`
	IsActive     bool              `json:"is_active"`
	Variants     []VariantResponse `json:"variants"`
	CreatedAt    int64             `json:"created_at"`
	UpdatedAt    int64             `json:"updated_at"`
}
