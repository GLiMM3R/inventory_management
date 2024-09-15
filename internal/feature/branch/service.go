package branch

import (
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/feature/user"
	"inverntory_management/internal/types"

	"github.com/google/uuid"
)

type BranchServiceImpl interface {
	GetAll(page, limit int, userCalims types.UserClaims, notSelf bool) ([]schema.Branch, int64, error)
	FindByID(branch_id string) (*schema.Branch, error)
	Create(dto BranchCreateDto) error
	Update(branch_id string, dto BranchUpdateDto) error
}

type branchService struct {
	branchRepo BranchRepositoryImpl
	userRepo   user.UserRepositoryImpl
}

func NewBranchService(branchRepo BranchRepositoryImpl, userRepo user.UserRepositoryImpl) BranchServiceImpl {
	return &branchService{branchRepo: branchRepo, userRepo: userRepo}
}

// GetAll implements BranchServiceImpl.
func (s *branchService) GetAll(page int, limit int, userClaims types.UserClaims, notSelf bool) ([]schema.Branch, int64, error) {
	user, err := s.userRepo.FindByID(userClaims.Subject)
	if err != nil {
		return nil, 0, err
	}

	branches, total, err := s.branchRepo.GetAll(page, limit, user.BranchID, notSelf)
	if err != nil {
		return nil, 0, err
	}

	return branches, total, nil
}

// Create implements BranchServiceImpl.
func (s *branchService) Create(dto BranchCreateDto) error {
	newBranch := &schema.Branch{
		BranchID: uuid.NewString(),
		Name:     dto.Name,
	}

	if err := s.branchRepo.Create(newBranch); err != nil {
		return err
	}

	return nil
}

// FindByID implements BranchServiceImpl.
func (s *branchService) FindByID(branch_id string) (*schema.Branch, error) {
	branch, err := s.branchRepo.FindByID(branch_id)
	if err != nil {
		return nil, err
	}

	return branch, nil
}

// Update implements BranchServiceImpl.
func (s *branchService) Update(branch_id string, dto BranchUpdateDto) error {
	// Find the existing branch by ID
	existingBranch, err := s.branchRepo.FindByID(branch_id)
	if err != nil {
		return err
	}

	// Update the fields of the existing branch with the new values from the DTO
	existingBranch.Name = dto.Name

	// Save the updated branch back to the repository
	if err := s.branchRepo.Update(existingBranch); err != nil {
		return err
	}

	return nil
}
