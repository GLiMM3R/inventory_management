package user

import (
	"errors"
	"fmt"

	"inverntory_management/internal/database/schema"
	"inverntory_management/internal/exception"
	custom "inverntory_management/pkg/errors"

	"gorm.io/gorm"
)

type UserRepositoryImpl interface {
	GetAll(page, limit int) ([]schema.User, int64, error)
	FindByID(user_id string) (*schema.User, error)
	FindByUsername(username string) (*schema.User, error)
	Create(user *schema.User) error
	Update(user *schema.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryImpl {
	return &userRepository{db: db}
}

// Create implements UserRepository.
func (r *userRepository) Create(user *schema.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}

	return nil
}

// FindByID implements UserRepository.
func (r *userRepository) FindByID(user_id string) (*schema.User, error) {
	var user *schema.User

	if err := r.db.First(&user, "user_id = ?", user_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, custom.NewDataNotFoundError("user not found")
		}
		return nil, custom.NewInternalServerError("internal server error")
	}

	return user, nil
}

// FindByUsername implements UserRepository.
func (r *userRepository) FindByUsername(username string) (*schema.User, error) {
	var user *schema.User

	if err := r.db.First(&user, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("here 1")
			return nil, custom.NewDataNotFoundError("user not found")
		}

		fmt.Println("here 2")

		return nil, custom.NewInternalServerError("internal server error")
	}
	fmt.Println("here 3")

	return user, nil
}

// GetAll implements UserRepository.
func (r *userRepository) GetAll(page int, limit int) ([]schema.User, int64, error) {
	var data []schema.User
	var total int64
	offset := (page - 1) * limit

	query := r.db.Model(&schema.User{})

	if err := query.Count(&total).Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

// Update implements UserRepository.
func (r *userRepository) Update(user *schema.User) error {
	if err := r.db.Save(&user).Error; err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}
