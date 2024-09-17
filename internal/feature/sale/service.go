package sale

import (
	"fmt"
	"inverntory_management/internal/database/schema"
	"sync"
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
	return &saleService{saleRepo: saleRepo}
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

func (s *saleService) Create(dto SaleCreateDto) error {
	errChan := make(chan error, len(dto.Items))
	var wg sync.WaitGroup

	count, err := s.saleRepo.Count()
	if err != nil {
		return err
	}

	orderNumber := fmt.Sprintf("ORD-%09d", count+1)
	saleDate := time.Now().Unix()

	for _, item := range dto.Items {
		wg.Add(1)
		go func(item Item) {
			defer wg.Done()
			newSale := &schema.Sale{
				SaleID:      uuid.NewString(),
				InventoryID: item.InventoryID,
				OrderNumber: orderNumber,
				Quantity:    item.Quantity,
				SaleDate:    saleDate,
			}

			if err := s.saleRepo.Create(newSale); err != nil {
				errChan <- err
			}
		}(item)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}
