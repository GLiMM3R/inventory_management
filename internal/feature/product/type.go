package product

import "inverntory_management/internal/feature/media"

type ProductCreateDto struct {
	Name        string                     `json:"name" validate:"required"`
	CategoryID  string                     `json:"category_id" validate:"required"`
	Description string                     `json:"description"`
	Images      []media.CreateMediaRequest `json:"images,omitempty"`
	Variants    []VariantCreateDto         `json:"variants"`
}

type VariantCreateDto struct {
	Price      float64              `json:"price" validate:"required"`
	Attributes []AttributeCreateDto `json:"attributes"`
}

type AttributeCreateDto struct {
	Attribute string `json:"attribute"`
	Value     string `json:"value"`
}

type ProductUpdateDto struct {
	Name        *string             `json:"name,omitempty" validate:"omitempty"`
	CategoryID  *string             `json:"category_id,omitempty" validate:"omitempty"`
	Description *string             `json:"description,omitempty" validate:"omitempty"`
	Images      *[]string           `json:"images,omitempty" validate:"omitempty"`
	Variants    *[]VariantUpdateDto `json:"variants,omitempty" validate:"omitempty"`
}

type VariantUpdateDto struct {
	VariantID  *string               `json:"variant_id,omitempty" validate:"omitempty"`
	Price      *float64              `json:"price,omitempty" validate:"omitempty"`
	Attributes *[]AttributeUpdateDto `json:"attributes,omitempty" validate:"omitempty"`
}

type AttributeUpdateDto struct {
	AttributeID *string `json:"attribute_id,omitempty" validate:"omitempty"`
	Attribute   *string `json:"attribute,omitempty" validate:"omitempty"`
	Value       *string `json:"value,omitempty" validate:"omitempty"`
}

type ProductListResponse struct {
	ProductID   string                `json:"product_id"`
	Name        string                `json:"name"`
	Images      []media.MediaResponse `json:"images"`
	Category    string                `json:"category"`
	Description string                `json:"description"`
	Variants    []VariantResponse     `json:"variants"`
	CreatedAt   int64                 `json:"created_at"`
	UpdatedAt   int64                 `json:"updated_at"`
}

type VariantResponse struct {
	VariantID  string              `json:"variant_id"`
	SKU        string              `json:"sku"`
	Price      float64             `json:"price"`
	Attributes []AttributeResponse `json:"attributes"`
}

type AttributeResponse struct {
	AttributeID string `json:"attribute_id"`
	Attribute   string `json:"attribute"`
	Value       string `json:"value"`
}

type ProductResponse struct {
	ProductID   string                `json:"product_id"`
	Name        string                `json:"name"`
	Images      []media.MediaResponse `json:"images"`
	CategoryID  string                `json:"category_id"`
	Category    string                `json:"category"`
	Description string                `json:"description"`
	Variants    []VariantResponse     `json:"variants"`
	CreatedAt   int64                 `json:"created_at"`
	UpdatedAt   int64                 `json:"updated_at"`
}
