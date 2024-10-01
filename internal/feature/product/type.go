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
