package variant

import "inverntory_management/internal/feature/media"

type VariantCreateDto struct {
	SKU          string                    `json:"sku" validate:"required"`
	Price        float64                   `json:"price" validate:"min=0"`
	Quantity     int                       `json:"quantity" validate:"min=0"`
	RestockLevel int                       `json:"restock_level" validate:"min=0"`
	Image        *media.CreateMediaRequest `json:"images,omitempty"`
	Attributes   []AttributeCreateDto      `json:"attributes"`
}

type AttributeCreateDto struct {
	Attribute string `json:"attribute"`
	Value     string `json:"value"`
}

type VariantUpdateDto struct {
	SKU          *string                   `json:"sku,omitempty" validate:"omitempty"`
	Price        *float64                  `json:"price,omitempty" validate:"min=0,omitempty"`
	Quantity     *int                      `json:"quantity,omitempty" validate:"min=0,omitempty"`
	RestockLevel *int                      `json:"restock_level,omitempty" validate:"min=0,omitempty"`
	Image        *media.CreateMediaRequest `json:"images,omitempty"`
	Attributes   *[]AttributeUpdateDto     `json:"attributes,omitempty" validate:"omitempty"`
}

type AttributeUpdateDto struct {
	AttributeID *string `json:"attribute_id,omitempty" validate:"omitempty"`
	Attribute   *string `json:"attribute,omitempty" validate:"omitempty"`
	Value       *string `json:"value,omitempty" validate:"omitempty"`
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
