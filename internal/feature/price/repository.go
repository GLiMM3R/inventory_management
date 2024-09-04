package price

import (
	"errors"
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/exception"

	"gorm.io/gorm"
)

type PriceRepositoryImpl interface {
	GetAll(page, limit int) ([]schema.Price, int64, error)
	FindByID(price_id string) (*schema.Price, error)
	Create(price *schema.Price) error
	Update(price *schema.Price) error
}

type priceRepository struct {
	db *gorm.DB
}

func NewPriceRepository(db *gorm.DB) PriceRepositoryImpl {
	return &priceRepository{db: db}
}

// Create implements PriceRepositoryImpl.
func (r *priceRepository) Create(price *schema.Price) error {
	if err := r.db.Create(&price).Error; err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}

// FindByID implements PriceRepositoryImpl.
func (r *priceRepository) FindByID(price_id string) (*schema.Price, error) {
	var price *schema.Price

	if err := r.db.First(&price, "price_id = ?", price_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, exception.ErrInternal
	}
	return price, nil
}

// GetAll implements PriceRepositoryImpl.
func (r *priceRepository) GetAll(page int, limit int) ([]schema.Price, int64, error) {
	var data []schema.Price
	var total int64
	offset := (page - 1) * limit

	query := r.db.Model(&schema.Price{})

	if err := query.Count(&total).Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

// Update implements PriceRepositoryImpl.
func (r *priceRepository) Update(price *schema.Price) error {
	if err := r.db.Save(&price).Error; err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}
