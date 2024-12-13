package user

import (
	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/exception"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl interface {
	GetAll(page, limit int) ([]schema.User, int64, error)
	FindByID(user_id string) (*schema.User, error)
	FindByUsername(username string) (*schema.User, error)
	Create(dto UserCreateDto) error
}

type userService struct {
	userRepo UserRepositoryImpl
}

func NewUserService(userRepo UserRepositoryImpl) UserServiceImpl {
	return &userService{userRepo: userRepo}
}

// FindByID implements Service.
func (s *userService) FindByID(user_id string) (*schema.User, error) {
	user, err := s.userRepo.FindByID(user_id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByUsername implements Service.
func (s *userService) FindByUsername(username string) (*schema.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetAll implements Service.
func (s *userService) GetAll(page int, limit int) ([]schema.User, int64, error) {
	users, total, err := s.userRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Create implements Service.
func (s *userService) Create(dto UserCreateDto) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return exception.NewAppError("HASH_ERROR", "Failed to hash password", err)
	}

	newUser := &schema.User{
		UserID:   uuid.NewString(),
		Email:    dto.Email,
		Username: dto.Username,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.Create(newUser); err != nil {
		return err
	}

	return nil
}
