package sale

import (
	"errors"
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/exception"

	"gorm.io/gorm"
)

type SaleRepositoryImpl interface {
	GetAll(page, limit int) ([]schema.Sale, int64, error)
	FindByID(sale_id string) (*schema.Sale, error)
	Create(sale *schema.Sale) error
	Update(sale *schema.Sale) error
	Count() (int64, error)
}

type saleRepository struct {
	db *gorm.DB
}

func NewSaleRepository(db *gorm.DB) SaleRepositoryImpl {
	return &saleRepository{db: db}
}

// Create implements PriceRepositoryImpl.
func (r *saleRepository) Create(sale *schema.Sale) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var existingInventory schema.Inventory

		if err := r.db.First(&existingInventory, "inventory_id = ?", sale.InventoryID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return exception.ErrNotFound
			}
			return exception.ErrInternal
		}

		if sale.Quantity > existingInventory.Quantity {
			if err := tx.Model(&schema.Inventory{}).Where("inventory_id = ?", sale.InventoryID).
				Update("status", "sold").Error; err != nil {
				return exception.ErrInternal
			}
			return exception.ErrInsufficientQuantity
		}

		quantityFloat := float64(sale.Quantity)
		sale.TotalPrice = quantityFloat * existingInventory.Price

		if err := tx.Model(&schema.Inventory{}).Where("inventory_id = ?", sale.InventoryID).UpdateColumn("quantity", gorm.Expr("quantity - ?", sale.Quantity)).Error; err != nil {
			return exception.ErrInternal
		}

		if err := tx.Create(&sale).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return exception.ErrDuplicateEntry
			}
			return exception.ErrInternal
		}
		return nil
	})
}

// FindByID implements PriceRepositoryImpl.
func (r *saleRepository) FindByID(sale_id string) (*schema.Sale, error) {
	var sale *schema.Sale

	if err := r.db.First(&sale, "sale_id = ?", sale_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, exception.ErrInternal
	}
	return sale, nil
}

// GetAll implements PriceRepositoryImpl.
func (r *saleRepository) GetAll(page int, limit int) ([]schema.Sale, int64, error) {
	var data []schema.Sale
	var total int64
	offset := (page - 1) * limit

	query := r.db.Model(&schema.Sale{})

	if err := query.Count(&total).Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

// Update implements PriceRepositoryImpl.
func (r *saleRepository) Update(sale *schema.Sale) error {
	if err := r.db.Save(&sale).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}

// Count implements SaleRepositoryImpl.
func (r *saleRepository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&schema.Sale{}).Count(&count).Error; err != nil {
		return 0, exception.ErrInternal
	}
	return count, nil
}
