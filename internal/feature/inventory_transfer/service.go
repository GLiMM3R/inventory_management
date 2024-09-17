package inventory_transfer

import (
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/feature/user"
	"inverntory_management/internal/types"
	"time"

	"github.com/google/uuid"
)

type InventoryTransferServiceImpl interface {
	GetAll(page, limit int, startDateUnix, endDateUnix int64, userClaims types.UserClaims) ([]InventoryTransferResponse, int64, error)
	FindByID(transfer_id string) (*schema.InventoryTransfer, error)
	Create(dto InventoryTransferCreateDto, userClaims types.UserClaims) error
}

type inventoryTransferService struct {
	inventoryRepo InventoryTransferRepositoryImpl
	userRepo      user.UserRepositoryImpl
}

func NewInventoryService(inventoryRepo InventoryTransferRepositoryImpl, userRepo user.UserRepositoryImpl) InventoryTransferServiceImpl {
	return &inventoryTransferService{inventoryRepo: inventoryRepo, userRepo: userRepo}
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
func (s *inventoryTransferService) GetAll(page, limit int, startDateUnix, endDateUnix int64, userClaims types.UserClaims) ([]InventoryTransferResponse, int64, error) {
	user, err := s.userRepo.FindByID(userClaims.Subject)
	if err != nil {
		return nil, 0, err
	}

	inventories, total, err := s.inventoryRepo.GetAll(page, limit, user.BranchID, startDateUnix, endDateUnix)
	if err != nil {
		return nil, 0, err
	}

	return inventories, total, nil
}

// Create implements InventoryServiceImpl.
func (s *inventoryTransferService) Create(dto InventoryTransferCreateDto, userClaims types.UserClaims) error {
	user, err := s.userRepo.FindByID(userClaims.Subject)
	if err != nil {
		return err
	}

	newTransfer := &schema.InventoryTransfer{
		TransferID:   uuid.NewString(),
		InventoryID:  dto.InventoryID,
		FromBranchID: user.BranchID,
		ToBranchID:   dto.ToBranchID,
		Quantity:     dto.Quantity,
		TransferDate: time.Now().Unix(),
	}

	if err := s.inventoryRepo.Create(newTransfer); err != nil {
		return err
	}

	return nil
}
