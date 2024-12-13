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
	"log"
	"path"
	"time"

	"github.com/google/uuid"
)

type VariantService interface {
	FindByID(variant_id string) (*VariantResponse, error)
	Create(variant_id string, request CreateVariantDTO) error
	Update(variant_id string, request UpdateVariantDTO) error
	Delete(variant_id string) error
}

type variantService struct {
	vairantRepo VariantRepository
	productRepo product.ProductRepositoryImpl
	s3Client    aws_service.S3Client
	cfg         config.Config
}

func NewVariantService(variantRepo VariantRepository, productRepo product.ProductRepositoryImpl, s3Client aws_service.S3Client, cfg config.Config) VariantService {
	return &variantService{vairantRepo: variantRepo, productRepo: productRepo, s3Client: s3Client, cfg: cfg}
}

// FindByID implements VariantService.
func (s *variantService) FindByID(variant_id string) (*VariantResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	variant, err := s.vairantRepo.FindByID(variant_id)
	if err != nil {
		return nil, err
	}

	response := &VariantResponse{
		VariantID:       variant.VariantID,
		SKU:             variant.SKU,
		VariantName:     variant.VariantName,
		AdditionalPrice: variant.AdditionalPrice,
		StockQuantity:   variant.StockQuantity,
		RestockLevel:    variant.RestockLevel,
		IsActive:        variant.IsActive,
		Status:          variant.Status,
		Attributes:      make([]AttributeResponse, len(variant.Attributes)),
		CreatedAt:       variant.CreatedAt,
		UpdatedAt:       variant.UpdatedAt,
	}

	if variant.Image != nil {
		result, err := s.s3Client.GetObject(ctx, s.cfg.AWS_BUCKET_NAME, variant.Image.Path, variant.Image.Type, int64(3600))
		if err != nil {
			log.Println(err.Error())
		}

		response.ImageURL = result.URL
	}

	for i, attr := range variant.Attributes {
		response.Attributes[i] = AttributeResponse{
			AttributeName:  attr.AttributeName,
			AttributeValue: attr.AttributeValue,
		}
	}

	return response, nil
}

// AddVariant implements ProductServiceImpl.
func (s *variantService) Create(product_id string, request CreateVariantDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product, err := s.productRepo.FindById(product_id)
	if err != nil {
		return err
	}

	newVariant := schema.ProductVariant{
		VariantID:       uuid.NewString(),
		ProductID:       product.ProductID,
		SKU:             request.SKU,
		AdditionalPrice: request.AdditionalPrice,
		StockQuantity:   request.StockQuantity,
		RestockLevel:    request.RestockLevel,
		Attributes:      make([]schema.Attribute, len(request.Attributes)),
	}

	if request.Image != nil {
		sourcePath := path.Join("tmp", request.Image.Name)
		destPath := path.Join("products", product.ProductID, request.Image.Name)

		if err := s.s3Client.CopyToFolder(ctx, s.cfg.AWS_BUCKET_NAME, sourcePath, destPath); err != nil {
			return err_response.NewInternalServerError()
		}

		if err := s.s3Client.DeleteObjects(ctx, config.AppConfig.AWS_BUCKET_NAME, []string{sourcePath}); err != nil {
			return err_response.NewInternalServerError()
		}

		newVariant.Image = &schema.Media{
			ID:             uuid.NewString(),
			Name:           request.Image.Name,
			Path:           destPath,
			Size:           request.Image.Size,
			Type:           request.Image.Type,
			CollectionType: request.Image.CollectionType,
		}
	}

	for idx, attribute := range request.Attributes {
		newVariant.Attributes[idx] = schema.Attribute{
			AttributeName:  attribute.AttributeName,
			AttributeValue: attribute.AttributeValue,
		}
	}

	if err := s.vairantRepo.Create(&newVariant); err != nil {
		return err
	}

	return nil
}

// UpdateVariant implements ProductServiceImpl.
func (s *variantService) Update(variant_id string, request UpdateVariantDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	vairant, err := s.vairantRepo.FindByID(variant_id)
	if err != nil {
		return err
	}

	if request.SKU != nil {
		vairant.SKU = *request.SKU
	}

	if request.AdditionalPrice != nil {
		vairant.AdditionalPrice = *request.AdditionalPrice
	}

	if request.StockQuantity != nil {
		vairant.StockQuantity = *request.StockQuantity
	}

	if request.RestockLevel != nil {
		vairant.RestockLevel = *request.RestockLevel
	}

	if request.IsActive != nil {
		vairant.IsActive = *request.IsActive
	}

	if request.Status != nil {
		vairant.Status = *request.Status
	}

	if request.Attributes != nil {
		for idx, attReq := range *request.Attributes {
			if *attReq.AttributeName == vairant.Attributes[idx].AttributeName {
				vairant.Attributes[idx].AttributeValue = *attReq.AttributeValue
			} else {
				newAttribute := schema.Attribute{
					AttributeName:  *attReq.AttributeName,
					AttributeValue: *attReq.AttributeValue,
				}

				vairant.Attributes = append(vairant.Attributes, newAttribute)
			}
		}
	}

	if request.Image != nil {
		if vairant.Image != nil {
			if err := s.s3Client.DeleteObjects(ctx, s.cfg.AWS_BUCKET_NAME, []string{vairant.Image.Path}); err != nil {
				return err_response.NewInternalServerError()
			}
		}

		sourcePath := path.Join("tmp", request.Image.Name)
		destPath := path.Join("products", vairant.ProductID, request.Image.Name)

		if err := s.s3Client.CopyToFolder(ctx, s.cfg.AWS_BUCKET_NAME, sourcePath, destPath); err != nil {
			return err_response.NewInternalServerError()
		}

		if err := s.s3Client.DeleteObjects(ctx, s.cfg.AWS_BUCKET_NAME, []string{sourcePath}); err != nil {
			return err_response.NewInternalServerError()
		}

		vairant.Image = &schema.Media{
			ID:             uuid.NewString(),
			Name:           request.Image.Name,
			Path:           destPath,
			Size:           request.Image.Size,
			Type:           request.Image.Type,
			CollectionType: request.Image.CollectionType,
		}
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
