package branch

import (
	"errors"
	"inverntory_management/internal/exception"

	"gorm.io/gorm"
)

type BranchRepositoryImpl interface {
	GetAll(page, limit int) ([]Branch, int64, error)
	FindByID(branch_id string) (*Branch, error)
	Create(branch *Branch) error
	Update(branch *Branch) error
}

type branchRepository struct {
	db *gorm.DB
}

func NewBranchRepository(db *gorm.DB) BranchRepositoryImpl {
	return &branchRepository{db: db}
}

// Create implements BranchRepository.
func (r *branchRepository) Create(branch *Branch) error {
	if err := r.db.Create(branch).Error; err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}

// GetAll implements BranchRepository.
func (r *branchRepository) GetAll(page int, limit int) ([]Branch, int64, error) {
	var data []Branch
	var total int64
	offset := (page - 1) * limit

	query := r.db.Model(&Branch{})

	if err := query.Count(&total).Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

// GetByID implements BranchRepository.
func (r *branchRepository) FindByID(branch_id string) (*Branch, error) {
	var branch *Branch

	if err := r.db.Where("branch_id = ?", branch_id).First(&branch).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, exception.ErrInternal
	}
	return branch, nil
}

// Update implements BranchRepository.
func (r *branchRepository) Update(branch *Branch) error {
	if err := r.db.Save(&branch).Error; err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}
