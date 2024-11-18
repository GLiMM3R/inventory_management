package category

import (
	"errors"
	"inverntory_management/internal/database/schema"
	custom "inverntory_management/pkg/errors"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl interface {
	FindAll(page int, limit int, parent_id string) ([]schema.Category, int64, error)
	Create(category *schema.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepositoryImpl {
	return &categoryRepository{db: db}
}

// Create implements CategoryRepositoryImpl.
func (r *categoryRepository) Create(category *schema.Category) error {
	if err := r.db.Create(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return custom.NewConflictError("Duplicate key")
		}
		return custom.NewInternalServerError()
	}
	return nil
}

// FindAll implements CategoryRepositoryImpl.
func (r *categoryRepository) FindAll(page int, limit int, parent_id string) ([]schema.Category, int64, error) {
	var data []schema.Category
	var total int64
	offset := (page - 1) * limit

	query := r.db.Model(&schema.Category{})

	if parent_id != "" {
		query = query.Where("parent_category_id = ?", parent_id)
	}

	if err := query.Count(&total).Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}
