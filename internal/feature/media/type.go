package media

type CreateMediaDTO struct {
	Name           string `json:"name"`
	Path           string `json:"path"`
	Type           string `json:"type"`
	Size           uint   `json:"size"`
	CollectionType string `json:"collection_type"`
}

type MediaResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Path           string `json:"path"`
	Type           string `json:"type"`
	Size           uint   `json:"size"`
	URL            string `json:"url"`
	CollectionType string `json:"collection_type"`
}
