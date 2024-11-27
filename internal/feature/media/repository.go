package media

import (
	"inverntory_management/internal/database/schema"
	error_response "inverntory_management/pkg/errors"

	"gorm.io/gorm"
)

type MediaRepository interface {
	Create(media *schema.Media) error
	GetByID(id int) (*schema.Media, error)
	GetAll(ids []string) ([]schema.Media, error)
	Update(media *schema.Media) error
	Delete(id int) error
}

type mediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) MediaRepository {
	return &mediaRepository{db: db}
}

// Create implements MediaRepository.
func (r *mediaRepository) Create(media *schema.Media) error {
	if err := r.db.Create(media).Error; err != nil {
		return error_response.NewInternalServerError()
	}
	return nil
}

// Delete implements MediaRepository.
func (r *mediaRepository) Delete(id int) error {
	panic("unimplemented")
}

// GetAll implements MediaRepository.
func (r *mediaRepository) GetAll(ids []string) ([]schema.Media, error) {
	var data []schema.Media

	query := r.db.Model(&schema.Media{})

	if len(ids) > 0 {
		query = query.Where("id IN ?", ids)
	}

	if err := query.Find(&data).Error; err != nil {
		return nil, error_response.NewInternalServerError()
	}

	return data, nil
}

// GetByID implements MediaRepository.
func (r *mediaRepository) GetByID(id int) (*schema.Media, error) {
	var data schema.Media
	if err := r.db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, error_response.NewInternalServerError()
	}
	return &data, nil
}

// Update implements MediaRepository.
func (r *mediaRepository) Update(media *schema.Media) error {
	if err := r.db.Save(media).Error; err != nil {
		return error_response.NewInternalServerError()
	}
	return nil
}
