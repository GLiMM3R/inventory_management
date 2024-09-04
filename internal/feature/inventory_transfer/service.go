package inventory_transfer

import (
	"inverntory_management/internal/database/schema"
	"time"

	"github.com/google/uuid"
)

type InventoryTransferServiceImpl interface {
	GetAll(page, limit int) ([]schema.InventoryTransfer, int64, error)
	FindByID(transfer_id string) (*schema.InventoryTransfer, error)
	Create(dto InventoryTransferCreateDto) error
}

type inventoryTransferService struct {
	inventoryRepo InventoryTransferRepositoryImpl
}

func NewInventoryService(inventoryRepo InventoryTransferRepositoryImpl) InventoryTransferServiceImpl {
	return &inventoryTransferService{inventoryRepo: inventoryRepo}
}

// FindByID implements InventoryServiceImpl.
func (s *inventoryTransferService) FindByID(inventory_id string) (*schema.InventoryTransfer, error) {
	inventory, err := s.inventoryRepo.FindByID(inventory_id)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

// GetAll implements InventoryServiceImpl.
func (s *inventoryTransferService) GetAll(page int, limit int) ([]schema.InventoryTransfer, int64, error) {
	inventories, total, err := s.inventoryRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	return inventories, total, nil
}

// Create implements InventoryServiceImpl.
func (s *inventoryTransferService) Create(dto InventoryTransferCreateDto) error {
	newTransfer := &schema.InventoryTransfer{
		TransferID:   uuid.NewString(),
		InventoryID:  dto.InventoryID,
		FromBranchID: dto.FromBranchID,
		ToBranchID:   dto.ToBranchID,
		Quantity:     dto.Quantity,
		TransferDate: time.Now().Unix(),
	}

	if err := s.inventoryRepo.Create(newTransfer); err != nil {
		return err
	}

	return nil
}
