package inventory

import (
	"errors"
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/exception"

	"gorm.io/gorm"
)

type InventoryRepositoryImpl interface {
	GetAll(page, limit int, branchID string) ([]schema.Inventory, int64, error)
	FindByID(inventory_id string) (*schema.Inventory, error)
	Create(inventory *schema.Inventory) error
	Update(inventory *schema.Inventory) error
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepositoryImpl {
	return &inventoryRepository{db: db}
}

// Create implements InventoryRepositoryImpl.
func (r *inventoryRepository) Create(inventory *schema.Inventory) error {
	if err := r.db.Create(&inventory).Error; err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}

// FindByID implements InventoryRepositoryImpl.
func (r *inventoryRepository) FindByID(inventory_id string) (*schema.Inventory, error) {
	var inventory *schema.Inventory

	if err := r.db.First(&inventory, "inventory_id = ?", inventory_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, exception.ErrInternal
	}
	return inventory, nil
}

// GetAll implements InventoryRepositoryImpl.
func (r *inventoryRepository) GetAll(page int, limit int, branchID string) ([]schema.Inventory, int64, error) {
	var data []schema.Inventory
	var total int64
	offset := (page - 1) * limit
	query := r.db.Model(&schema.Inventory{})
	if err := query.Where("fk_branch_id = ?", branchID).Count(&total).Limit(limit).Offset(offset).Order("created_at DESC").Find(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

// Update implements InventoryRepositoryImpl.
func (r *inventoryRepository) Update(inventory *schema.Inventory) error {
	if err := r.db.Save(&inventory).Error; err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}
