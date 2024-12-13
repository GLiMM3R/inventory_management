package variant

import (
	"errors"
	"inverntory_management/internal/database/schema"
	err_response "inverntory_management/pkg/errors"
	"log"

	"gorm.io/gorm"
)

type VariantRepository interface {
	//variant
	FindByID(variant_id string) (*schema.ProductVariant, error)
	Create(variant *schema.ProductVariant) error
	Update(variant *schema.ProductVariant) error
	Delete(variant_id string) error
}

type variantRepository struct {
	db *gorm.DB
}

func NewVariantRepository(db *gorm.DB) VariantRepository {
	return &variantRepository{db: db}
}

// FindByID implements VariantRepository.
func (r *variantRepository) FindByID(variant_id string) (*schema.ProductVariant, error) {
	var data schema.ProductVariant

	if err := r.db.Preload("Attributes").First(&data, "variant_id = ?", variant_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("%s", err.Error())
			return nil, err_response.NewNotFoundError("Variant not found!")
		}
		log.Printf("%s", err.Error())
		return nil, err_response.NewInternalServerError()
	}

	return &data, nil
}

// CreateVariant implements ProductRepositoryImpl.
func (r *variantRepository) Create(variant *schema.ProductVariant) error {
	if err := r.db.Create(&variant).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("%s", err.Error())
			return err_response.NewConflictError("Duplicate key")
		}
		log.Printf("%s", err.Error())
		return err_response.NewInternalServerError()
	}
	return nil
}

// DeleteVariant implements ProductRepositoryImpl.
func (r *variantRepository) Delete(variant_id string) error {
	if err := r.db.Delete(&schema.ProductVariant{}, "variant_id = ?", variant_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("%s", err.Error())
			return err_response.NewNotFoundError("Variant not found!")
		}
		log.Printf("%s", err.Error())
		return err_response.NewInternalServerError()
	}
	return nil
}

// UpdateVariant implements ProductRepositoryImpl.
func (r *variantRepository) Update(variant *schema.ProductVariant) error {
	if err := r.db.Save(&variant).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("%s", err.Error())
			return err_response.NewConflictError("Duplicate key")
		}
		log.Printf("%s", err.Error())
		return err_response.NewInternalServerError()
	}
	return nil
}
