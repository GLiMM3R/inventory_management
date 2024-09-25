package inventory_transfer

import (
	"errors"
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/exception"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InventoryTransferRepositoryImpl interface {
	GetAll(page, limit int, branchID string, startDateUnix, endDateUnix int64) ([]InventoryTransferResponse, int64, error)
	FindByID(transfer_id string) (*schema.InventoryTransfer, error)
	Create(transfer *schema.InventoryTransfer) error
	Update(transfer *schema.InventoryTransfer) error
}

type inventoryTransferRepository struct {
	db *gorm.DB
}

func NewInventoryTransferRepository(db *gorm.DB) InventoryTransferRepositoryImpl {
	return &inventoryTransferRepository{db: db}
}

// Create implements PriceRepositoryImpl.
func (r *inventoryTransferRepository) Create(transfer *schema.InventoryTransfer) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var inventory *schema.Inventory

		if err := tx.First(&inventory, "inventory_id = ?", transfer.InventoryID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return exception.ErrNotFound
			}
			return exception.ErrInternal
		}

		if inventory.Quantity < transfer.Quantity {
			if err := tx.Model(&schema.Inventory{}).Where("inventory_id = ?", inventory.InventoryID).
				Update("status", "sold").Error; err != nil {
				return exception.ErrInternal
			}
			return exception.ErrInsufficientQuantity
		}

		var existingInventory *schema.Inventory

		tx.Where("fk_branch_id = ?", transfer.ToBranchID).Where("fk_product_id = ?", inventory.ProductID).
			First(&existingInventory)

		if existingInventory.ID != 0 {
			existingInventory.Quantity = existingInventory.Quantity + transfer.Quantity
			if err := tx.Model(&schema.Inventory{}).Where("inventory_id = ?", inventory.InventoryID).
				Update("quantity", existingInventory.Quantity).Error; err != nil {
				if errors.Is(err, gorm.ErrDuplicatedKey) {
					return exception.ErrDuplicateEntry
				}
				return exception.ErrInternal
			}
		} else {
			newInventory := &schema.Inventory{
				InventoryID:  uuid.NewString(),
				BranchID:     transfer.ToBranchID,
				ProductID:    transfer.Inventory.ProductID,
				Quantity:     transfer.Quantity,
				RestockLevel: 0,
				IsActive:     true,
			}

			if err := tx.Create(&newInventory).Error; err != nil {
				if errors.Is(err, gorm.ErrDuplicatedKey) {
					return exception.ErrDuplicateEntry
				}
				return exception.ErrInternal
			}
		}

		inventory.Quantity -= transfer.Quantity

		if err := tx.Model(&schema.Inventory{}).Where("inventory_id = ?", inventory.InventoryID).
			Update("quantity", inventory.Quantity).Error; err != nil {
			return exception.ErrInternal
		}

		if err := tx.Create(&transfer).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return exception.ErrDuplicateEntry
			}
			return exception.ErrInternal
		}
		return nil
	})
}

// FindByID implements PriceRepositoryImpl.
func (r *inventoryTransferRepository) FindByID(price_id string) (*schema.InventoryTransfer, error) {
	var transfer *schema.InventoryTransfer

	if err := r.db.First(&transfer, "transfer_id = ?", price_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, exception.ErrInternal
	}
	return transfer, nil
}

// GetAll implements PriceRepositoryImpl.
func (r *inventoryTransferRepository) GetAll(page int, limit int, branchID string, startDateUnix, endDateUnix int64) ([]InventoryTransferResponse, int64, error) {
	var data []InventoryTransferResponse
	var total int64
	offset := (page - 1) * limit

	query := r.db.Model(&schema.InventoryTransfer{})

	if err := query.Select("inventory_transfers.transfer_id as transfer_id", "inventory_transfers.fk_inventory_id as inventory_id",
		"inventories.name as inventory_name", "inventory_transfers.fk_to_branch_id as to_branch_id", "branches.name as to_branch",
		"inventory_transfers.quantity as quantity", "inventory_transfers.transfer_date as transfer_date").
		Where("fk_from_branch_id = ? AND transfer_date >= ? AND transfer_date <= ?", branchID, startDateUnix, endDateUnix).
		Joins("left join inventories on inventories.inventory_id = inventory_transfers.fk_inventory_id").
		Joins("left join branches on branches.branch_id = inventory_transfers.fk_to_branch_id").
		Count(&total).Limit(limit).Offset(offset).Scan(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

// Update implements PriceRepositoryImpl.
func (r *inventoryTransferRepository) Update(transfer *schema.InventoryTransfer) error {
	if err := r.db.Save(&transfer).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}
