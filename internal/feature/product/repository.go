package product

import (
	"errors"
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/exception"

	"gorm.io/gorm"
)

type ProductRepositoryImpl interface {
	Create(product *schema.Product) error
	Update(product *schema.Product) error
	Delete(id uint) error
	FindById(product_id string) (*schema.Product, error)
	FindAll(page int, limit int) ([]schema.Product, int64, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryImpl {
	return &productRepository{db: db}
}

// Create implements ProductRepositoryImpl.
func (r *productRepository) Create(product *schema.Product) error {
	if err := r.db.Create(&product).Error; err != nil {
		return exception.ErrInternal
	}
	return nil
}

// Delete implements ProductRepositoryImpl.
func (r *productRepository) Delete(id uint) error {
	panic("unimplemented")
}

// FindAll implements ProductRepositoryImpl.
func (r *productRepository) FindAll(page int, limit int) ([]schema.Product, int64, error) {
	var data []schema.Product
	var total int64
	offset := (page - 1) * limit

	query := r.db.Model(&schema.Price{})

	if err := query.Count(&total).Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

// FindById implements ProductRepositoryImpl.
func (r *productRepository) FindById(product_id string) (*schema.Product, error) {
	var product *schema.Product

	if err := r.db.First(&product, "product_id = ?", product_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, exception.ErrInternal
	}
	return product, nil
}

// Update implements ProductRepositoryImpl.
func (r *productRepository) Update(product *schema.Product) error {
	if err := r.db.Save(&product).Error; err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}
