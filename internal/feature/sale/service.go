package sale

import (
	"inverntory_management/internal/database/schema"
	"time"

	"github.com/google/uuid"
)

type SaleServiceImpl interface {
	GetAll(page, limit int) ([]schema.Sale, int64, error)
	FindByID(sale_id string) (*schema.Sale, error)
	Create(dto SaleCreateDto) error
	// Update(price_id string, dto PriceUpdateDto) error
}

type saleService struct {
	saleRepo SaleRepositoryImpl
}

func NewSaleService(saleRepo SaleRepositoryImpl) SaleServiceImpl {
	return &saleService{
		saleRepo: saleRepo,
	}
}

// FindByID implements PriceServiceImpl.
func (s *saleService) FindByID(sale_id string) (*schema.Sale, error) {
	price, err := s.saleRepo.FindByID(sale_id)
	if err != nil {
		return nil, err
	}

	return price, nil
}

// GetAll implements PriceServiceImpl.
func (s *saleService) GetAll(page int, limit int) ([]schema.Sale, int64, error) {
	sales, total, err := s.saleRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	return sales, total, nil
}

// Create implements PriceServiceImpl.
func (s *saleService) Create(dto SaleCreateDto) error {
	saleID := uuid.NewString()

	for _, item := range dto.Items {
		newSale := &schema.Sale{
			SaleID:      saleID,
			InventoryID: item.InventoryID,
			Quantity:    item.Quantity,
			SaleDate:    time.Now().Unix(),
		}

		if err := s.saleRepo.Create(newSale); err != nil {
			return err
		}
	}

	// if err := s.saleRepo.Create(sales); err != nil {
	// 	return err
	// }

	return nil
}
