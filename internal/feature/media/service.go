package media

import (
	"inverntory_management/internal/database/schema"
	aws_service "inverntory_management/pkg/aws"
	"path/filepath"

	"github.com/google/uuid"
)

type MediaService interface {
	FindAll(ids []string) ([]schema.Media, error)
	FindByID(id int) (*schema.Media, error)
	Create(media *schema.Media) (*schema.Media, error)
	Update(id int, media *schema.Media) (*schema.Media, error)
	Delete(id int) error
}

type mediaService struct {
	mediaRepository MediaRepository
	s3client        aws_service.S3Client
}

func NewMediaService(mediaRepository MediaRepository, s3client aws_service.S3Client) MediaService {
	return &mediaService{
		mediaRepository: mediaRepository,
		s3client:        s3client,
	}
}

// Create implements MediaService.
func (s *mediaService) Create(media *schema.Media) (*schema.Media, error) {
	newMediaID := uuid.NewString()
	ext := filepath.Ext(media.Name)
	fileName := newMediaID + "." + ext

	newMedia := &schema.Media{
		ID:             newMediaID,
		Name:           fileName,
		Type:           media.Type,
		Size:           media.Size,
		Path:           media.Path,
		CollectionType: media.CollectionType,
	}

	err := s.mediaRepository.Create(newMedia)
	if err != nil {
		return nil, err
	}

	return newMedia, nil
}

// Delete implements MediaService.
func (s *mediaService) Delete(id int) error {
	panic("unimplemented")
}

// FindAll implements MediaService.
func (s *mediaService) FindAll(ids []string) ([]schema.Media, error) {
	medias, err := s.mediaRepository.GetAll(ids)
	if err != nil {
		return nil, err
	}

	return medias, nil
}

// FindByID implements MediaService.
func (s *mediaService) FindByID(id int) (*schema.Media, error) {
	media, err := s.mediaRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return media, nil
}

// Update implements MediaService.
func (s *mediaService) Update(id int, media *schema.Media) (*schema.Media, error) {
	find, err := s.mediaRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	find.Name = media.Name
	find.Type = media.Type
	find.Size = media.Size
	find.Path = media.Path
	find.CollectionType = media.CollectionType

	err = s.mediaRepository.Update(media)
	if err != nil {
		return nil, err
	}

	return media, nil
}
