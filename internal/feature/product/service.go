package product

import (
	"context"
	"encoding/json"
	"fmt"
	"inverntory_management/config"
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/feature/media"
	"inverntory_management/internal/utils"
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
	Create(product ProductCreateDto) error
	Update(product_id string, product ProductUpdateDto) error
}

type productService struct {
	productRepo ProductRepositoryImpl
	mediaRepo   media.MediaRepository
	s3Client    aws_service.S3Client
}

func NewProductService(productRepo ProductRepositoryImpl, mediaRepo media.MediaRepository, s3Client aws_service.S3Client) ProductServiceImpl {
	return &productService{productRepo: productRepo, mediaRepo: mediaRepo, s3Client: s3Client}
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
			ProductID:   product.ProductID,
			Name:        product.Name,
			Category:    product.Category.Name,
			Description: product.Description,
			Variants:    make([]VariantResponse, len(product.Variants)),
			Images:      make([]media.MediaResponse, len(product.Images)),
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}

		for j, variant := range product.Variants {
			response[i].Variants[j] = VariantResponse{
				VariantID: variant.VariantID,
				SKU:       variant.SKU,
			}
		}

		if len(product.Images) > 0 {
			for j, image := range product.Images {
				result, err := s.s3Client.GetObject(ctx, config.AppConfig.AWS_BUCKET_NAME, image.Media.FilePath, int64(3600))
				if err != nil {
					log.Println(err)
				}

				response[i].Images[j] = media.MediaResponse{
					MediaID:     image.MediaID,
					FileURL:     result.URL,
					FileName:    image.Media.FileName,
					FilePath:    image.Media.FilePath,
					FileType:    image.Media.FileType,
					FileSize:    image.Media.FileSize,
					MediaType:   image.Media.MediaType,
					Description: image.Media.Description,
					CreatedAt:   image.Media.CreatedAt,
					UpdatedAt:   image.Media.UpdatedAt,
				}
			}
		}
	}

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
		ProductID:   product.ProductID,
		Name:        product.Name,
		CategoryID:  product.CategoryID,
		Category:    product.Category.Name,
		Description: product.Description,
		Variants:    make([]VariantResponse, len(product.Variants)),
		Images:      make([]media.MediaResponse, len(product.Images)),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
	for j, variant := range product.Variants {
		response.Variants[j] = VariantResponse{
			VariantID:    variant.VariantID,
			SKU:          variant.SKU,
			Price:        variant.Price,
			Quantity:     variant.Quantity,
			RestockLevel: variant.RestockLevel,
			IsActive:     variant.IsActive,
			Status:       variant.Status,
			CreatedAt:    variant.CreatedAt,
			UpdatedAt:    variant.UpdatedAt,
			Attributes:   make([]AttributeResponse, len(variant.Attributes)),
		}

		if variant.Image != nil {
			result, err := s.s3Client.GetObject(ctx, config.AppConfig.AWS_BUCKET_NAME, variant.Image.FilePath, int64(3600))
			if err != nil {
				log.Println(err.Error())
			}

			response.Variants[j].Image = media.MediaResponse{
				MediaID:     variant.Image.MediaID,
				FileURL:     result.URL,
				FileName:    variant.Image.FileName,
				FilePath:    variant.Image.FilePath,
				FileType:    variant.Image.FileType,
				FileSize:    variant.Image.FileSize,
				MediaType:   variant.Image.MediaType,
				Description: variant.Image.Description,
				CreatedAt:   variant.Image.CreatedAt,
				UpdatedAt:   variant.Image.UpdatedAt,
			}
		}

		for k, attribute := range variant.Attributes {
			response.Variants[j].Attributes[k] = AttributeResponse{
				AttributeID: attribute.AttributeID,
				Attribute:   attribute.Attribute,
				Value:       attribute.Value,
			}
		}
	}

	if len(product.Images) > 0 {
		for j, image := range product.Images {
			result, err := s.s3Client.GetObject(ctx, config.AppConfig.AWS_BUCKET_NAME, image.Media.FilePath, int64(3600))
			if err != nil {
				log.Println(err.Error())
			}

			response.Images[j] = media.MediaResponse{
				MediaID:     image.MediaID,
				FileURL:     result.URL,
				FileName:    image.Media.FileName,
				FilePath:    image.Media.FilePath,
				FileType:    image.Media.FileType,
				FileSize:    image.Media.FileSize,
				MediaType:   image.Media.MediaType,
				Description: image.Media.Description,
				CreatedAt:   image.Media.CreatedAt,
				UpdatedAt:   image.Media.UpdatedAt,
			}
		}
	}

	return &response, nil
}

// Update implements ProductServiceImpl.
func (s *productService) Update(product_id string, request ProductUpdateDto) error {
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

	if request.Description != nil {
		product.Description = *request.Description
	}

	for _, variantReq := range *request.Variants {
		//check update variant
		if variantReq.VariantID != nil {
			//loop update vairant
			for i, v := range product.Variants {
				//check if req has id
				if v.VariantID == *variantReq.VariantID {
					//check update attribute
					if variantReq.Attributes != nil {
						for _, attributeReq := range *variantReq.Attributes {
							if attributeReq.AttributeID != nil {
								for j, a := range product.Variants[i].Attributes {
									if a.AttributeID == *attributeReq.AttributeID {
										if attributeReq.Attribute != nil {
											fmt.Println(attributeReq.Attribute)
											product.Variants[i].Attributes[j] = schema.Attribute{
												ID:          a.ID,
												VariantID:   a.VariantID,
												AttributeID: *attributeReq.AttributeID,
												Attribute:   *attributeReq.Attribute,
												Value:       *attributeReq.Value,
											}
										}
									}
								}
							}
							//  else {
							// 	newAttribute := schema.Attribute{
							// 		AttributeID: uuid.NewString(),
							// 		VariantID:   *variantReq.VariantID,
							// 		Attribute:   *attributeReq.Attribute,
							// 		Value:       *attributeReq.Value,
							// 	}
							// 	product.Variants[i].Attributes = append(product.Variants[i].Attributes, newAttribute)
							// }
						}
					}
				}
			}
		} else {
			newVariantID := uuid.NewString()
			SKU := product.Name

			newVariant := schema.Variant{
				VariantID:  newVariantID,
				Attributes: make([]schema.Attribute, len(*variantReq.Attributes)),
			}

			product.Variants = append(product.Variants, newVariant)

			for i, attribute := range *variantReq.Attributes {
				newAttribute := schema.Attribute{
					AttributeID: uuid.NewString(),
					VariantID:   newVariantID,
					Attribute:   *attribute.Attribute,
					Value:       *attribute.Value,
				}
				SKU += " " + newAttribute.Value

				product.Variants[len(product.Variants)-1].Attributes[i] = newAttribute
			}

			newVariant.SKU = utils.GenerateSKU(SKU, 14, "SKU-", 3)
		}
	}

	productJSON, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(productJSON))

	err = s.productRepo.Update(product)
	if err != nil {
		return err
	}

	return nil
}

// Create implements ProductRepositoryImpl.
func (s *productService) Create(product ProductCreateDto) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newProduct := schema.Product{
		ProductID:   uuid.NewString(),
		Name:        product.Name,
		CategoryID:  product.CategoryID,
		Description: product.Description,
		Images:      make([]schema.ProductMedia, len(product.Images)),
	}

	for idx, image := range product.Images {
		newMediaID := uuid.NewString()
		sourcePath := path.Join("tmp", image.FileName)
		destPath := path.Join("products", newProduct.ProductID, image.FileName)

		if err := s.s3Client.CopyToFolder(ctx, config.AppConfig.AWS_BUCKET_NAME, sourcePath, destPath); err != nil {
			return err_response.NewInternalServerError()
		}

		if err := s.s3Client.DeleteObjects(ctx, config.AppConfig.AWS_BUCKET_NAME, []string{sourcePath}); err != nil {
			return err_response.NewInternalServerError()
		}

		newProduct.Images[idx] = schema.ProductMedia{
			MediaID:   newMediaID,
			ProductID: newProduct.ProductID,
			Media: schema.Media{
				MediaID:     newMediaID,
				FileName:    image.FileName,
				FileType:    image.FileType,
				FileSize:    image.FileSize,
				FilePath:    destPath,
				MediaType:   image.MediaType,
				Description: image.Description,
			},
		}
	}

	for _, variant := range product.Variants {
		newVariantID := uuid.NewString()
		SKU := product.Name

		newVariant := schema.Variant{
			VariantID:  newVariantID,
			Attributes: make([]schema.Attribute, len(variant.Attributes)),
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
	}

	if err := s.productRepo.Create(&newProduct); err != nil {
		return err
	}

	return nil
}
