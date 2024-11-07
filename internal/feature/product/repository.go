package product

import (
	"errors"
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/exception"
	custom "inverntory_management/pkg/errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepositoryImpl interface {
	//product
	Create(product *schema.Product) error
	Update(product *schema.Product) error
	Delete(product_id string) error
	FindById(product_id string) (*schema.Product, error)
	FindAll(page int, limit int) ([]schema.Product, int64, error)

	//variant
	CreateVariant(variant *schema.Variant) error
	UpdateVariant(variant *schema.Variant) error
	DeleteVariant(product_id string) error
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
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return custom.NewConflictError("Duplicate key")
		}
		return custom.NewInternalServerError()
	}
	return nil
}

// Delete implements ProductRepositoryImpl.
func (r *productRepository) Delete(product_id string) error {
	panic("unimplemented")
}

// FindAll implements ProductRepositoryImpl.
func (r *productRepository) FindAll(page int, limit int) ([]schema.Product, int64, error) {
	var data []schema.Product
	var total int64
	offset := (page - 1) * limit

	query := r.db.Model(&schema.Product{})

	if err := query.Preload("Category").Count(&total).Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

// FindById implements ProductRepositoryImpl.
func (r *productRepository) FindById(product_id string) (*schema.Product, error) {
	var product *schema.Product

	if err := r.db.Preload("Variants").Preload("Variants.Attributes").Preload("Variants.Price").Preload(clause.Associations).First(&product, "product_id = ?", product_id).Error; err != nil {
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
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}

// CreateVariant implements ProductRepositoryImpl.
func (r *productRepository) CreateVariant(variant *schema.Variant) error {
	if err := r.db.Create(&variant).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return exception.ErrDuplicateEntry
		}
		return err
	}
	return nil
}

// DeleteVariant implements ProductRepositoryImpl.
func (r *productRepository) DeleteVariant(product_id string) error {
	panic("unimplemented")
}

// UpdateVariant implements ProductRepositoryImpl.
func (r *productRepository) UpdateVariant(variant *schema.Variant) error {
	if err := r.db.Save(&variant).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}
