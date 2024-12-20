package category

import (
	"inverntory_management/internal/database/schema"

	"github.com/google/uuid"
)

type CategoryServiceImpl interface {
	GetAll(page, limit int, parent_id string) ([]schema.Category, int64, error)
	Create(category CategoryRequest) error
}

type categoryService struct {
	repo CategoryRepositoryImpl
}

func NewCategoryService(repo CategoryRepositoryImpl) CategoryServiceImpl {
	return &categoryService{repo: repo}
}

// Create implements CategoryServiceImpl.
func (s *categoryService) Create(category CategoryRequest) error {
	var newCategory schema.Category

	newCategory.CategoryID = uuid.NewString()
	newCategory.Name = category.Name
	newCategory.Level = category.Level

	if category.ParentCategoryID != nil {
		newCategory.ParentCategoryID = category.ParentCategoryID
	}

	err := s.repo.Create(&newCategory)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements CategoryRepositoryImpl.
func (s *categoryService) GetAll(page int, limit int, parent_id string) ([]schema.Category, int64, error) {
	categories, total, err := s.repo.FindAll(page, limit, parent_id)
	if err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}
