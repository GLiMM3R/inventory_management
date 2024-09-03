package inventory

import (
	"fmt"
	"inverntory_management/internal/database/schema"

	"github.com/google/uuid"
)

type InventoryServiceImpl interface {
	GetAll(page, limit int) ([]schema.Inventory, int64, error)
	FindByID(inventory_id string) (*schema.Inventory, error)
	Create(dto InventoryCreateDto) error
	Update(inventory_id string, dto InventoryUpdateDto) error
}

type inventoryService struct {
	inventoryRepo InventoryRepositoryImpl
}

func NewInventoryService(inventoryRepo InventoryRepositoryImpl) InventoryServiceImpl {
	return &inventoryService{inventoryRepo: inventoryRepo}
}

// FindByID implements InventoryServiceImpl.
func (s *inventoryService) FindByID(inventory_id string) (*schema.Inventory, error) {
	inventory, err := s.inventoryRepo.FindByID(inventory_id)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

// GetAll implements InventoryServiceImpl.
func (s *inventoryService) GetAll(page int, limit int) ([]schema.Inventory, int64, error) {
	inventories, total, err := s.inventoryRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	return inventories, total, nil
}

// Create implements InventoryServiceImpl.
func (s *inventoryService) Create(dto InventoryCreateDto) error {
	newInventory := &schema.Inventory{
		InventoryID: uuid.NewString(),
		BranchID:    dto.BranchID,
		Name:        dto.Name,
		SKU:         dto.SKU,
		Quantity:    dto.Quantity,
		Status:      ACTIVE,
	}

	if err := s.inventoryRepo.Create(newInventory); err != nil {
		return err
	}

	return nil
}

// Update implements InventoryServiceImpl.
func (s *inventoryService) Update(inventory_id string, dto InventoryUpdateDto) error {
	existingInventory, err := s.inventoryRepo.FindByID(inventory_id)
	if err != nil {
		return err
	}

	existingInventory.Name = *dto.Name
	existingInventory.SKU = *dto.SKU
	existingInventory.Quantity = *dto.Quantity
	existingInventory.Status = *dto.Status

	fmt.Println(existingInventory)

	// if err := s.inventoryRepo.Update(existingInventory); err != nil {
	// 	return err
	// }

	return nil
}
