package schema

type Media struct {
	ID             string `json:"id" gorm:"primaryKey;unique;column:id"`
	Name           string `json:"name" gorm:"column:name"`
	Path           string `json:"path" gorm:"column:path"`
	Type           string `json:"type" gorm:"column:type"`
	Size           uint   `json:"size" gorm:"column:size"`
	CollectionType string `json:"collection_type" gorm:"column:collection_type"`
}
