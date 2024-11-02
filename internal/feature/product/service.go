package product

import (
	"fmt"
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/utils"
	"time"

	"github.com/google/uuid"
)

type ProductServiceImpl interface {
	//product
	FindAll(page, limit int) ([]ProductResponse, int64, error)
	FindByID(product_id string) (*schema.Product, error)
	Create(product ProductCreateDto) error
	Update(product schema.Product) error

	//variant
	AddVariant(product_id string, variant VariantCreateDto) error
}

type productService struct {
	productRepo ProductRepositoryImpl
}

func NewProductService(productRepo ProductRepositoryImpl) ProductServiceImpl {
	return &productService{productRepo: productRepo}
}

// FindAll implements ProductServiceImpl.
func (s *productService) FindAll(page int, limit int) ([]ProductResponse, int64, error) {
	products, total, err := s.productRepo.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	response := make([]ProductResponse, len(products))

	for i, product := range products {
		fmt.Println(product.Images)
		response[i] = ProductResponse{
			ProductID:   product.ProductID,
			Name:        product.Name,
			Images:      product.Images,
			Category:    product.Category.Name,
			Description: product.Description,
			Variants:    make([]VariantResponse, len(product.Variants)),
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}
	}

	return response, total, nil
}

// FindByID implements ProductServiceImpl.
func (s *productService) FindByID(product_id string) (*schema.Product, error) {
	product, err := s.productRepo.FindById(product_id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// Update implements ProductServiceImpl.
func (s *productService) Update(product schema.Product) error {
	panic("unimplemented")
}

// Create implements ProductRepositoryImpl.
func (s *productService) Create(product ProductCreateDto) error {
	newProduct := schema.Product{
		ProductID:   uuid.NewString(),
		Name:        product.Name,
		CategoryID:  product.CategoryID,
		Description: product.Description,
		Images:      product.Images,
	}

	for _, variant := range product.Variants {
		newVariantID := uuid.NewString()
		SKU := product.Name

		newVariant := schema.Variant{
			VariantID:  newVariantID,
			Attributes: make([]schema.Attribute, len(variant.Attributes)),
			Price:      make([]schema.PriceHistory, 1),
		}

		for idx, attribute := range variant.Attributes {
			SKU += " " + attribute.Value
			newVariant.Attributes[idx] = schema.Attribute{
				AttributeID: uuid.NewString(),
				Attribute:   attribute.Attribute,
				Value:       attribute.Value,
			}
		}

		newVariant.SKU = utils.GenerateSKU(SKU, 14, "SKU-", 3)

		newProduct.Variants = append(newProduct.Variants, newVariant)

		newVariant.Price[0] = schema.PriceHistory{
			PriceID:       uuid.NewString(),
			NewPrice:      variant.Price,
			OldPrice:      0,
			EffectiveDate: time.Now().Unix(),
		}
	}

	if err := s.productRepo.Create(&newProduct); err != nil {
		return err
	}

	return nil
}

// AddVariant implements ProductServiceImpl.
func (s *productService) AddVariant(product_id string, variant VariantCreateDto) error {
	product, err := s.productRepo.FindById(product_id)
	if err != nil {
		return err
	}

	SKU := product.Name

	newVariant := schema.Variant{
		VariantID:  uuid.NewString(),
		ProductID:  product.ProductID,
		Attributes: make([]schema.Attribute, len(variant.Attributes)),
		Price:      make([]schema.PriceHistory, 1),
	}

	for idx, attribute := range variant.Attributes {
		SKU += " " + attribute.Value
		newVariant.Attributes[idx] = schema.Attribute{
			AttributeID: uuid.NewString(),
			Attribute:   attribute.Attribute,
			Value:       attribute.Value,
		}
	}

	newVariant.SKU = utils.GenerateSKU(SKU, 14, "SKU-", 3)

	newVariant.Price[0] = schema.PriceHistory{
		PriceID:       uuid.NewString(),
		NewPrice:      variant.Price,
		OldPrice:      0,
		EffectiveDate: time.Now().Unix(),
	}

	if err := s.productRepo.CreateVariant(&newVariant); err != nil {
		return err
	}

	return nil
}
