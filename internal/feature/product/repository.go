package product

import (
	"errors"
	"inverntory_management/internal/database/schema"
	err_response "inverntory_management/pkg/errors"
	"log"

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
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryImpl {
	return &productRepository{db: db}
}

// Create implements ProductRepositoryImpl.
func (r *productRepository) Create(product *schema.Product) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&product).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return err_response.NewConflictError("Duplicate key")
			}
			return err_response.NewInternalServerError()
		}
		return nil
	})
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

	if err := query.Preload("Category").Preload("Thumbail").Count(&total).Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("%s", err.Error())
			return nil, 0, err_response.NewNotFoundError("No products found!")
		}
		return nil, 0, err_response.NewInternalServerError()
	}

	return data, total, nil
}

// FindById implements ProductRepositoryImpl.
func (r *productRepository) FindById(product_id string) (*schema.Product, error) {
	var product *schema.Product

	if err := r.db.Preload("Category").Preload("Variants").Preload("Variants.Attributes").Preload("Variants.Image").Preload("Thumbail").Preload(clause.Associations).First(&product, "product_id = ?", product_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("%s", err.Error())
			return nil, err_response.NewNotFoundError("Product not found!")
		}
		log.Printf("%s", err.Error())
		return nil, err_response.NewInternalServerError()
	}
	return product, nil
}

// Update implements ProductRepositoryImpl.
func (r *productRepository) Update(product *schema.Product) error {
	// if err := r.db.Model(&product).Select("name", "category_id", "description").Updates(product).Error; err != nil {
	if err := r.db.Model(&product).Updates(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("%s", err.Error())
			return err_response.NewConflictError("Duplicate key")
		}
		log.Printf("%s", err.Error())
		return err_response.NewInternalServerError()
	}
	return nil
}
