package branch

import "github.com/google/uuid"

type BranchServiceImpl interface {
	GetAll(page, limit int) ([]Branch, int64, error)
	FindByID(branch_id string) (*Branch, error)
	Create(dto BranchCreateDto) error
	Update(branch_id string, dto BranchUpdateDto) error
}

type branchService struct {
	branchRepo BranchRepositoryImpl
}

func NewBranchService(branchRepo BranchRepositoryImpl) BranchServiceImpl {
	return &branchService{branchRepo: branchRepo}
}

// GetAll implements BranchServiceImpl.
func (s *branchService) GetAll(page int, limit int) ([]Branch, int64, error) {
	users, total, err := s.branchRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Create implements BranchServiceImpl.
func (s *branchService) Create(dto BranchCreateDto) error {
	newBranch := &Branch{
		BranchID: uuid.NewString(),
		Name:     dto.Name,
	}

	if err := s.branchRepo.Create(newBranch); err != nil {
		return err
	}

	return nil
}

// FindByID implements BranchServiceImpl.
func (s *branchService) FindByID(branch_id string) (*Branch, error) {
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
