package inventory

import (
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/feature/user"
	"inverntory_management/internal/types"
	"inverntory_management/internal/utils"

	"github.com/google/uuid"
)

type InventoryServiceImpl interface {
	GetAll(page, limit int, userClaims *types.UserClaims) ([]schema.Inventory, int64, error)
	FindByID(inventory_id string) (*schema.Inventory, error)
	Create(dto InventoryCreateDto, userClaims *types.UserClaims) error
	Update(inventory_id string, dto InventoryUpdateDto) error
}

type inventoryService struct {
	inventoryRepo InventoryRepositoryImpl
	userRepo      user.UserRepositoryImpl
}

func NewInventoryService(inventoryRepo InventoryRepositoryImpl, userRepo user.UserRepositoryImpl) InventoryServiceImpl {
	return &inventoryService{inventoryRepo: inventoryRepo, userRepo: userRepo}
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
func (s *inventoryService) GetAll(page int, limit int, userClaims *types.UserClaims) ([]schema.Inventory, int64, error) {
	user, err := s.userRepo.FindByID(userClaims.Subject)
	if err != nil {
		return nil, 0, err
	}

	inventories, total, err := s.inventoryRepo.GetAll(page, limit, user.BranchID)
	if err != nil {
		return nil, 0, err
	}

	return inventories, total, nil
}

// Create implements InventoryServiceImpl.
func (s *inventoryService) Create(dto InventoryCreateDto, userClaims *types.UserClaims) error {
	user, err := s.userRepo.FindByID(userClaims.Subject)
	if err != nil {
		return err
	}

	newInventory := &schema.Inventory{
		InventoryID: uuid.NewString(),
		BranchID:    user.BranchID,
		Name:        dto.Name,
		SKU:         utils.GenerateSKU(dto.Name, 12, "SKU-"),
		Quantity:    dto.Quantity,
		Price:       dto.Price,
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

	if dto.Name != nil {
		existingInventory.Name = *dto.Name
	}

	if dto.Quantity != nil {
		existingInventory.Quantity = *dto.Quantity
	}

	if dto.Price != nil {
		existingInventory.Price = *dto.Price
	}

	if dto.Status != nil {
		existingInventory.Status = *dto.Status
	}

	if err := s.inventoryRepo.Update(existingInventory); err != nil {
		return err
	}

	return nil
}
