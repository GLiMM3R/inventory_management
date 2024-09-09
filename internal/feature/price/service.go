package price

import (
	"inverntory_management/internal/database/schema"

	"github.com/google/uuid"
)

type PriceServiceImpl interface {
	GetAll(inventoryID string, page, limit int) ([]schema.Price, int64, error)
	FindByID(inventory_id string) (*schema.Price, error)
	Create(dto PriceCreateDto) error
	Update(price_id string, dto PriceUpdateDto) error
}

type priceService struct {
	priceRepo PriceRepositoryImpl
}

func NewPriceService(priceRepo PriceRepositoryImpl) PriceServiceImpl {
	return &priceService{
		priceRepo: priceRepo,
	}
}

// FindByID implements PriceServiceImpl.
func (s *priceService) FindByID(price_id string) (*schema.Price, error) {
	price, err := s.priceRepo.FindByID(price_id)
	if err != nil {
		return nil, err
	}

	return price, nil
}

// GetAll implements PriceServiceImpl.
func (s *priceService) GetAll(inventoryID string, page int, limit int) ([]schema.Price, int64, error) {
	prices, total, err := s.priceRepo.GetAll(inventoryID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return prices, total, nil
}

// Create implements PriceServiceImpl.
func (s *priceService) Create(dto PriceCreateDto) error {
	newPrice := &schema.Price{
		PriceID:       uuid.NewString(),
		InventoryID:   dto.InventoryID,
		Price:         dto.Price,
		EffectiveDate: dto.EffectiveDate,
	}

	if err := s.priceRepo.Create(newPrice); err != nil {
		return err
	}

	return nil
}

// Update implements PriceServiceImpl.
func (s *priceService) Update(price_id string, dto PriceUpdateDto) error {
	existingPrice, err := s.priceRepo.FindByID(price_id)
	if err != nil {
		return err
	}

	if dto.Price != nil {
		existingPrice.Price = *dto.Price
	}

	if dto.EffectiveDate != nil {
		existingPrice.EffectiveDate = *dto.EffectiveDate
	}

	if err := s.priceRepo.Update(existingPrice); err != nil {
		return err
	}

	return nil
}
