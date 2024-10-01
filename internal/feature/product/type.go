package product

type ProductCreateDto struct {
	Name        string             `json:"name" validate:"required"`
	CategoryID  string             `json:"category_id" validate:"required"`
	Description string             `json:"description"`
	Variants    []VariantCreateDto `json:"variants"`
}

type VariantCreateDto struct {
	Price      float64              `json:"price" validate:"required"`
	Attributes []AttributeCreateDto `json:"attributes"`
}

type AttributeCreateDto struct {
	Attribute string `json:"attribute"`
	Value     string `json:"value"`
}

type ProductResponse struct {
	ProductID   string            `json:"product_id"`
	Name        string            `json:"name"`
	Category    string            `json:"category"`
	Description string            `json:"description"`
	Variants    []VariantResponse `json:"variants"`
	CreatedAt   int64             `json:"created_at"`
	UpdatedAt   int64             `json:"updated_at"`
}

type VariantResponse struct {
	SKU        string              `json:"sku"`
	Price      float64             `json:"price"`
	Attributes []AttributeResponse `json:"attributes"`
}

type AttributeResponse struct {
	Attribute string `json:"attribute"`
	Value     string `json:"value"`
}
