package variant

import "inverntory_management/internal/feature/media"

type CreateVariantDto struct {
	SKU          string                    `json:"sku" validate:"required"`
	Price        float64                   `json:"price" validate:"min=0"`
	Quantity     int                       `json:"quantity" validate:"min=0"`
	RestockLevel int                       `json:"restock_level" validate:"min=0"`
	Image        *media.CreateMediaRequest `json:"images,omitempty"`
	Attributes   []CreateAttributeDto      `json:"attributes"`
}

type CreateAttributeDto struct {
	Attribute string `json:"attribute"`
	Value     string `json:"value"`
}

type UpdateVariantDto struct {
	SKU          *string                   `json:"sku,omitempty" validate:"omitempty"`
	Price        *float64                  `json:"price,omitempty" validate:"omitempty"`
	Quantity     *int                      `json:"quantity,omitempty" validate:"omitempty"`
	RestockLevel *int                      `json:"restock_level,omitempty" validate:"omitempty"`
	Image        *media.CreateMediaRequest `json:"images,omitempty"`
	Attributes   *[]UpdateAttributeDto     `json:"attributes,omitempty" validate:"omitempty"`
}

type UpdateAttributeDto struct {
	AttributeID *string `json:"attribute_id,omitempty" validate:"omitempty"`
	Attribute   *string `json:"attribute,omitempty" validate:"omitempty"`
	Value       *string `json:"value,omitempty" validate:"omitempty"`
}

type VariantResponse struct {
	VariantID    string              `json:"variant_id"`
	SKU          string              `json:"sku"`
	Price        float64             `json:"price"`
	Quantity     int                 `json:"quantity"`
	RestockLevel int                 `json:"restock_level"`
	IsActive     bool                `json:"is_active"`
	Status       string              `json:"status"`
	CreatedAt    int64               `json:"created_at"`
	UpdatedAt    int64               `json:"updated_at"`
	Image        media.MediaResponse `json:"image"`
	Attributes   []AttributeResponse `json:"attributes"`
}

type AttributeResponse struct {
	AttributeID string `json:"attribute_id"`
	Attribute   string `json:"attribute"`
	Value       string `json:"value"`
}
