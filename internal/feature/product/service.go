package product

import (
	"context"
	"encoding/json"
	"fmt"
	"inverntory_management/config"
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/feature/media"
	aws_service "inverntory_management/pkg/aws"
	err_response "inverntory_management/pkg/errors"
	"log"
	"path"
	"time"

	"github.com/google/uuid"
)

type ProductServiceImpl interface {
	//product
	FindAll(page, limit int) ([]ProductListResponse, int64, error)
	FindByID(product_id string) (*ProductResponse, error)
	Create(product CreateProductDTO) error
	Update(product_id string, product UpdateProductDTO) error
}

type productService struct {
	productRepo ProductRepositoryImpl
	mediaRepo   media.MediaRepository
	s3Client    aws_service.S3Client
	cfg         config.Config
}

func NewProductService(productRepo ProductRepositoryImpl, mediaRepo media.MediaRepository, s3Client aws_service.S3Client, cfg config.Config) ProductServiceImpl {
	return &productService{productRepo: productRepo, mediaRepo: mediaRepo, s3Client: s3Client, cfg: cfg}
}

// FindAll implements ProductServiceImpl.
func (s *productService) FindAll(page int, limit int) ([]ProductListResponse, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	products, total, err := s.productRepo.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	response := make([]ProductListResponse, len(products))

	for i, product := range products {
		response[i] = ProductListResponse{
			ProductID:    product.ProductID,
			Name:         product.Name,
			CategoryName: product.Category.Name,
			BasePrice:    product.BasePrice,
			Description:  product.Description,
			IsActive:     product.IsActive,
			CreatedAt:    product.CreatedAt,
			UpdatedAt:    product.UpdatedAt,
		}

		if product.ThumbnailID != nil {
			result, err := s.s3Client.GetObject(ctx, s.cfg.AWS_BUCKET_NAME, product.Thumbnail.Path, int64(3600))
			if err != nil {
				log.Println(err)
				continue
			}

			response[i].ThumbnailURL = result.URL
		}
	}

	productJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return nil, 0, err
	}
	fmt.Println(string(productJSON))

	return response, total, nil
}

// FindByID implements ProductServiceImpl.
func (s *productService) FindByID(product_id string) (*ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product, err := s.productRepo.FindById(product_id)
	if err != nil {
		return nil, err
	}

	response := ProductResponse{
		ProductID:    product.ProductID,
		Name:         product.Name,
		CategoryID:   product.CategoryID,
		CategoryName: product.Category.Name,
		BasePrice:    product.BasePrice,
		Description:  product.Description,
		Variants:     make([]VariantResponse, len(product.Variants)),
		IsActive:     product.IsActive,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}

	if product.ThumbnailID != nil {
		result, err := s.s3Client.GetObject(ctx, s.cfg.AWS_BUCKET_NAME, product.Thumbnail.Path, int64(3600))
		if err != nil {
			log.Println(err.Error())
		}

		response.ThumbnailURL = result.URL
	}

	for i, variant := range product.Variants {
		response.Variants[i] = VariantResponse{
			VariantID:       variant.VariantID,
			SKU:             variant.SKU,
			AdditionalPrice: variant.AdditionalPrice,
			StockQuantity:   variant.StockQuantity,
			RestockLevel:    variant.RestockLevel,
			IsActive:        variant.IsActive,
			Status:          variant.Status,
			CreatedAt:       variant.CreatedAt,
			UpdatedAt:       variant.UpdatedAt,
			Attributes:      make([]AttributeResponse, len(variant.Attributes)),
		}

		if variant.Image != nil {
			result, err := s.s3Client.GetObject(ctx, s.cfg.AWS_BUCKET_NAME, variant.Image.Path, int64(3600))
			if err != nil {
				log.Println(err.Error())
				continue
			}

			response.Variants[i].ImageURL = result.URL
		}

		for j, attribute := range variant.Attributes {
			response.Variants[i].Attributes[j] = AttributeResponse{
				AttributeName:  attribute.AttributeName,
				AttributeValue: attribute.AttributeValue,
			}
		}
	}

	return &response, nil
}

// Update implements ProductServiceImpl.
func (s *productService) Update(product_id string, request UpdateProductDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product, err := s.productRepo.FindById(product_id)
	if err != nil {
		return err
	}

	if request.CategoryID != nil {
		product.CategoryID = *request.CategoryID
	}

	if request.Name != nil {
		product.Name = *request.Name
	}

	if request.BasePrice != nil {
		product.BasePrice = *request.BasePrice
	}

	if request.Description != nil {
		product.Description = *request.Description
	}

	if request.IsActive != nil {
		product.IsActive = *request.IsActive
	}

	if request.Thumbnail != nil {
		if product.Thumbnail.Path != "" {
			_ = s.s3Client.DeleteObjects(ctx, s.cfg.AWS_BUCKET_NAME, []string{product.Thumbnail.Path})
		}

		sourcePath := path.Join("tmp", request.Thumbnail.Name)
		destPath := path.Join("products", product.ProductID, request.Thumbnail.Name)

		if err := s.s3Client.CopyToFolder(ctx, s.cfg.AWS_BUCKET_NAME, sourcePath, destPath); err != nil {
			return err_response.NewInternalServerError()
		}

		if err := s.s3Client.DeleteObjects(ctx, s.cfg.AWS_BUCKET_NAME, []string{sourcePath}); err != nil {
			return err_response.NewInternalServerError()
		}

		product.Thumbnail = &schema.Media{
			Name:           request.Thumbnail.Name,
			Path:           destPath,
			Type:           request.Thumbnail.Type,
			Size:           request.Thumbnail.Size,
			CollectionType: "thumbnail",
		}

	}

	err = s.productRepo.Update(product)
	if err != nil {
		return err
	}

	return nil
}

// Create implements ProductRepositoryImpl.
func (s *productService) Create(request CreateProductDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newProduct := schema.Product{
		ProductID:   uuid.NewString(),
		Name:        request.Name,
		CategoryID:  request.CategoryID,
		BasePrice:   request.BasePrice,
		Description: request.Description,
	}

	if request.Thumbnail != nil {
		newMediaID := uuid.NewString()
		sourcePath := path.Join("tmp", request.Thumbnail.Name)
		destPath := path.Join("products", newProduct.ProductID, request.Thumbnail.Name)

		if err := s.s3Client.CopyToFolder(ctx, s.cfg.AWS_BUCKET_NAME, sourcePath, destPath); err != nil {
			return err_response.NewInternalServerError()
		}

		if err := s.s3Client.DeleteObjects(ctx, s.cfg.AWS_BUCKET_NAME, []string{sourcePath}); err != nil {
			return err_response.NewInternalServerError()
		}

		newProduct.Thumbnail = &schema.Media{
			ID:             newMediaID,
			Name:           request.Thumbnail.Name,
			Path:           destPath,
			Type:           request.Thumbnail.Type,
			Size:           request.Thumbnail.Size,
			CollectionType: "thumbnail",
		}
	}

	for _, variantReq := range request.Variants {
		newVariantID := uuid.NewString()

		newVariant := &schema.ProductVariant{
			VariantID:       newVariantID,
			SKU:             variantReq.SKU,
			VariantName:     variantReq.VariantName,
			AdditionalPrice: variantReq.AdditionalPrice,
			StockQuantity:   variantReq.StockQuantity,
			RestockLevel:    variantReq.RestockLevel,
			Attributes:      make([]schema.Attribute, len(variantReq.Attributes)),
		}

		for idx, attribute := range variantReq.Attributes {
			newVariant.Attributes[idx] = schema.Attribute{
				VariantID:      newVariantID,
				AttributeName:  attribute.AttributeName,
				AttributeValue: attribute.AttributeValue,
			}
		}

		newProduct.Variants = append(newProduct.Variants, *newVariant)
	}

	if err := s.productRepo.Create(&newProduct); err != nil {
		return err
	}

	return nil
}
