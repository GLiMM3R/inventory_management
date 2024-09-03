package user

import (
	"errors"

	"inverntory_management/internal/exception"

	"gorm.io/gorm"
)

type UserRepositoryImpl interface {
	GetAll(page, limit int) ([]User, int64, error)
	FindByID(user_id string) (*User, error)
	FindByUsername(username string) (*User, error)
	Create(user *User) error
	Update(user *User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryImpl {
	return &userRepository{db: db}
}

// Create implements UserRepository.
func (r *userRepository) Create(user *User) error {
	if err := r.db.Create(user).Error; err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}

	return nil
}

// FindByID implements UserRepository.
func (r *userRepository) FindByID(user_id string) (*User, error) {
	var user *User

	if err := r.db.Where("user_id = ?", user_id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, exception.ErrInternal
	}

	return user, nil
}

// FindByUsername implements UserRepository.
func (r *userRepository) FindByUsername(username string) (*User, error) {
	var user *User

	if err := r.db.First(&user, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, exception.ErrInternal
	}

	return user, nil
}

// GetAll implements UserRepository.
func (r *userRepository) GetAll(page int, limit int) ([]User, int64, error) {
	var data []User
	var total int64
	offset := (page - 1) * limit

	query := r.db.Model(&User{})

	if err := query.Count(&total).Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

// Update implements UserRepository.
func (r *userRepository) Update(user *User) error {
	if err := r.db.Save(user).Error; err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return exception.ErrDuplicateEntry
		}
		return exception.ErrInternal
	}
	return nil
}
