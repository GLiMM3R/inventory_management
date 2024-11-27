package files

import (
	"context"
	"fmt"
	"inverntory_management/config"
	aws_service "inverntory_management/pkg/aws"
	custom "inverntory_management/pkg/errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type FileServiceImpl interface {
	UploadMultiFiles(fileType string, files []*multipart.FileHeader) ([]string, error)
	UploadFile(fileType string, file *multipart.FileHeader) (string, error)
	ReadFile(directory, fileName string) ([]byte, error)
	GeneratePresignPutObject(request PutObjectRequest) (*PutObjectResponse, error)
}

type fileService struct {
	s3Client aws_service.S3Client
}

func NewFileService(s3Client aws_service.S3Client) FileServiceImpl {
	return &fileService{s3Client: s3Client}
}

func (f *fileService) UploadMultiFiles(fileType string, files []*multipart.FileHeader) ([]string, error) {
	var fileNames []string

	fmt.Println(files[0].Filename)

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return nil, custom.NewInternalServerError()
		}
		defer src.Close()

		fileName := uuid.NewString() + filepath.Ext(file.Filename)
		fileNames = append(fileNames, fileName)

		dstPath := filepath.Join("uploads", fileType, fileName)
		dst, err := os.Create(dstPath)
		if err != nil {
			return nil, custom.NewInternalServerError()
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return nil, custom.NewInternalServerError()
		}
	}

	return fileNames, nil
}

// UploadFile implements FileServiceImpl.
func (f *fileService) UploadFile(fileType string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", custom.NewInternalServerError()
	}
	defer src.Close()

	fileName := uuid.NewString() + filepath.Ext(file.Filename)

	dstPath := filepath.Join("uploads", fileType, fileName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", custom.NewInternalServerError()
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", custom.NewInternalServerError()

	}

	return fileName, nil
}

func (f *fileService) ReadFile(directory, fileName string) ([]byte, error) {
	filePath := filepath.Join("uploads", directory, fileName)

	fmt.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, custom.NewInternalServerError()
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, custom.NewInternalServerError()
	}

	return fileData, nil
}

func (s *fileService) GeneratePresignPutObject(request PutObjectRequest) (*PutObjectResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fileName := uuid.NewString() + filepath.Ext(request.FileName)
	filePath := filepath.Join("tmp", fileName)

	res, err := s.s3Client.PutObject(ctx, config.AppConfig.AWS_BUCKET_NAME, filePath, int64(3600))
	if err != nil {
		return nil, custom.NewInternalServerError()
	}

	return &PutObjectResponse{
		URL:      res.URL,
		FileName: fileName,
		FilePath: filePath,
	}, nil
}
