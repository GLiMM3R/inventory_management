package schema

type Category struct {
	ID               uint      `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	CategoryID       string    `json:"category_id" gorm:"primaryKey;unique;column:category_id"`
	ParentCategoryID *string   `json:"parent_category_id" gorm:"column:parent_category_id;null"`
	Category         *Category `json:"-" gorm:"foreignKey:parent_category_id;references:category_id"`
	Name             string    `json:"name" gorm:"column:name"`
	Image            *string   `json:"image" gorm:"column:image"`
	CreatedAt        int64     `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt        int64     `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`

	// Reference to the parent category
}
