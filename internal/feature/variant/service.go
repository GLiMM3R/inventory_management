package variant

import (
	"context"
	"encoding/json"
	"fmt"
	"inverntory_management/config"
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/feature/product"
	aws_service "inverntory_management/pkg/aws"
	err_response "inverntory_management/pkg/errors"
	"path"
	"time"

	"github.com/google/uuid"
)

type VariantService interface {
	//variant
	FindByID(variant_id string) (*schema.Variant, error)
	Create(variant_id string, variant CreateVariantDto) error
	Update(variant_id string, variant UpdateVariantDto) error
	Delete(variant_id string) error
}

type variantService struct {
	vairantRepo VariantRepository
	productRepo product.ProductRepositoryImpl
	s3Client    aws_service.S3Client
}

func NewProductService(variantRepo VariantRepository, productRepo product.ProductRepositoryImpl, s3Client aws_service.S3Client) VariantService {
	return &variantService{vairantRepo: variantRepo, productRepo: productRepo, s3Client: s3Client}
}

// FindByID implements VariantService.
func (s *variantService) FindByID(variant_id string) (*schema.Variant, error) {
	variant, err := s.vairantRepo.FindByID(variant_id)
	if err != nil {
		return nil, err
	}

	return variant, nil
}

// AddVariant implements ProductServiceImpl.
func (s *variantService) Create(product_id string, req CreateVariantDto) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product, err := s.productRepo.FindById(product_id)
	if err != nil {
		return err
	}

	SKU := product.Name

	newVariant := schema.Variant{
		VariantID:    uuid.NewString(),
		ProductID:    product.ProductID,
		SKU:          req.SKU,
		Price:        req.Price,
		Quantity:     req.Quantity,
		RestockLevel: req.RestockLevel,
		Attributes:   make([]schema.Attribute, len(req.Attributes)),
	}

	if req.Image != nil {
		newMediaID := uuid.NewString()
		sourcePath := path.Join("tmp", req.Image.FileName)
		destPath := path.Join("products", product.ProductID, req.Image.FileName)

		if err := s.s3Client.CopyToFolder(ctx, config.AppConfig.AWS_BUCKET_NAME, sourcePath, destPath); err != nil {
			return err_response.NewInternalServerError()
		}

		if err := s.s3Client.DeleteObjects(ctx, config.AppConfig.AWS_BUCKET_NAME, []string{sourcePath}); err != nil {
			return err_response.NewInternalServerError()
		}

		newVariant.Image = &schema.Media{
			MediaID:     newMediaID,
			FileName:    req.Image.FileName,
			FileType:    req.Image.FileType,
			FileSize:    req.Image.FileSize,
			FilePath:    destPath,
			MediaType:   req.Image.MediaType,
			Description: req.Image.Description,
		}
	}

	for idx, attribute := range req.Attributes {
		SKU += " " + attribute.Value
		newVariant.Attributes[idx] = schema.Attribute{
			AttributeID: uuid.NewString(),
			Attribute:   attribute.Attribute,
			Value:       attribute.Value,
		}
	}

	if err := s.vairantRepo.Create(&newVariant); err != nil {
		return err
	}

	return nil
}

// UpdateVariant implements ProductServiceImpl.
func (s *variantService) Update(variant_id string, req UpdateVariantDto) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	vairant, err := s.vairantRepo.FindByID(variant_id)
	if err != nil {
		return err
	}

	if req.SKU != nil {
		vairant.SKU = *req.SKU
	}

	if req.Price != nil {
		vairant.Price = *req.Price
	}

	if req.Quantity != nil {
		vairant.Quantity = *req.Quantity
	}

	if req.RestockLevel != nil {
		vairant.RestockLevel = *req.RestockLevel
	}

	if req.Attributes != nil {
		for idx, attribute := range *req.Attributes {
			if attribute.AttributeID != nil {
				vairant.Attributes[idx].Attribute = *attribute.Attribute
				vairant.Attributes[idx].Value = *attribute.Value
			} else {
				newAttribute := schema.Attribute{
					AttributeID: uuid.NewString(),
					Attribute:   *attribute.Attribute,
					Value:       *attribute.Value,
				}
				vairant.Attributes = append(vairant.Attributes, newAttribute)
			}
		}
	}

	if req.Image != nil {
		fmt.Println("here=>")
		if vairant.Image != nil {
			if err := s.s3Client.DeleteObjects(ctx, config.AppConfig.AWS_BUCKET_NAME, []string{vairant.Image.FilePath}); err != nil {
				return err_response.NewInternalServerError()
			}
		}

		newMediaID := uuid.NewString()
		sourcePath := path.Join("tmp", req.Image.FileName)
		destPath := path.Join("products", vairant.ProductID, req.Image.FileName)

		if err := s.s3Client.CopyToFolder(ctx, config.AppConfig.AWS_BUCKET_NAME, sourcePath, destPath); err != nil {
			return err_response.NewInternalServerError()
		}

		if err := s.s3Client.DeleteObjects(ctx, config.AppConfig.AWS_BUCKET_NAME, []string{sourcePath}); err != nil {
			return err_response.NewInternalServerError()
		}

		vairant.Image = &schema.Media{
			MediaID:     newMediaID,
			FileName:    req.Image.FileName,
			FileType:    req.Image.FileType,
			FileSize:    req.Image.FileSize,
			FilePath:    destPath,
			MediaType:   req.Image.MediaType,
			Description: req.Image.Description,
		}
		vairant.ImageID = &newMediaID
	}

	productJSON, err := json.MarshalIndent(vairant, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(productJSON))

	if err := s.vairantRepo.Update(vairant); err != nil {
		return err
	}

	return nil
}

// Delete implements VariantService.
func (s *variantService) Delete(variant_id string) error {
	if err := s.vairantRepo.Delete(variant_id); err != nil {
		return err
	}

	return nil
}
