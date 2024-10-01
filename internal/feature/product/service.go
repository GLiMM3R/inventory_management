package product

import (
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/utils"
	"time"

	"github.com/google/uuid"
)

type ProductServiceImpl interface {
	FindAll(page, limit int) ([]schema.Product, int64, error)
	FindByID(product_id string) (*schema.Product, error)
	Create(product ProductCreateDto) error
	Update(product schema.Product) error
}

type productService struct {
	productRepo ProductRepositoryImpl
}

func NewProductService(productRepo ProductRepositoryImpl) ProductServiceImpl {
	return &productService{productRepo: productRepo}
}

// FindAll implements ProductServiceImpl.
func (s *productService) FindAll(page int, limit int) ([]schema.Product, int64, error) {
	panic("unimplemented")
}

// FindByID implements ProductServiceImpl.
func (s *productService) FindByID(product_id string) (*schema.Product, error) {
	panic("unimplemented")
}

// Update implements ProductServiceImpl.
func (s *productService) Update(product schema.Product) error {
	panic("unimplemented")
}

// Create implements ProductRepositoryImpl.
func (s *productService) Create(product ProductCreateDto) error {
	newProduct := schema.Product{
		Name:        product.Name,
		CategoryID:  product.CategoryID,
		Description: product.Description,
	}

	for _, variant := range product.Variants {
		newVariantID := uuid.NewString()

		newProduct.Variants = append(newProduct.Variants, schema.Variant{
			VariantID:  newVariantID,
			SKU:        utils.GenerateSKU(product.Name, 10, "SKU"),
			Price:      make([]schema.PriceHistory, 0),
			Attributes: make([]schema.Attribute, len(variant.Attributes)),
		})

		for idx, attribute := range variant.Attributes {
			newProduct.Variants[idx].Attributes = append(newProduct.Variants[idx].Attributes, schema.Attribute{
				AttributeID: uuid.NewString(),
				Attribute:   attribute.Attribute,
				Value:       attribute.Value,
			})
		}

		newPrice := schema.PriceHistory{
			PriceID:       uuid.NewString(),
			VariantID:     newVariantID,
			NewPrice:      variant.Price,
			OldPrice:      0,
			EffectiveDate: time.Now().Unix(),
		}

		newProduct.Variants[len(newProduct.Variants)-1].Price = append(newProduct.Variants[len(newProduct.Variants)-1].Price, newPrice)
	}

	if err := s.productRepo.Create(&newProduct); err != nil {
		return err
	}

	return nil
}
