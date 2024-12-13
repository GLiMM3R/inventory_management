package category

type CategoryRequest struct {
	Name             string  `json:"name"`
	ParentCategoryID *string `json:"parent_category_id,omitempty"`
	Level            int     `json:"level"`
}
