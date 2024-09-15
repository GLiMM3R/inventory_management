package branch

import (
	"errors"
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/exception"

	"gorm.io/gorm"
)

type BranchRepositoryImpl interface {
	GetAll(page, limit int, branchID string, notSelf bool) ([]schema.Branch, int64, error)
	FindByID(branch_id string) (*schema.Branch, error)
	Create(branch *schema.Branch) error
	Update(branch *schema.Branch) error
}

type branchRepository struct {
	db *gorm.DB
}

func NewBranchRepository(db *gorm.DB) BranchRepositoryImpl {
	return &branchRepository{db: db}
}

// Create implements BranchRepository.
func (r *branchRepository) Create(branch *schema.Branch) error {
	if err := r.db.Create(&branch).Error; err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}

// GetAll implements BranchRepository.
func (r *branchRepository) GetAll(page int, limit int, branchID string, notSelf bool) ([]schema.Branch, int64, error) {
	var data []schema.Branch
	var total int64
	offset := (page - 1) * limit

	query := r.db.Model(&schema.Branch{})

	if notSelf {
		query = query.Where("branch_id <> ?", branchID)
	}

	if err := query.Count(&total).Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

// GetByID implements BranchRepository.
func (r *branchRepository) FindByID(branch_id string) (*schema.Branch, error) {
	var branch *schema.Branch

	if err := r.db.Preload("Inventories").First(&branch, "branch_id = ?", branch_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, exception.ErrInternal
	}
	return branch, nil
}

// Update implements BranchRepository.
func (r *branchRepository) Update(branch *schema.Branch) error {
	if err := r.db.Save(&branch).Error; err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}
