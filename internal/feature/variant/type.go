package variant

import "inverntory_management/internal/feature/media"

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

type UpdateVariantDTO struct {
	SKU             *string               `json:"sku,omitempty"`
	VariantName     *string               `json:"variant_name,omitempty"`
	AdditionalPrice *float64              `json:"additional_price,omitempty" validate:"min=0"`
	StockQuantity   *int                  `json:"stock_quantity,omitempty" validate:"min=0"`
	RestockLevel    *int                  `json:"restock_level,omitempty" validate:"min=0"`
	IsActive        *bool                 `json:"is_active,omitempty"`
	Status          *string               `json:"status,omitempty"`
	Image           *media.CreateMediaDTO `json:"image,omitempty"`
	Attributes      *[]UpdateAttributeDTO `json:"attributes,omitempty"`
}

type UpdateAttributeDTO struct {
	AttributeName  *string `json:"attribute_name,omitempty"`
	AttributeValue *string `json:"attribute_value,omitempty"`
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
