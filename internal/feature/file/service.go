package files

import (
	"fmt"
	custom "inverntory_management/pkg/errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type FileServiceImpl interface {
	UploadMultiFiles(fileType string, files []*multipart.FileHeader) ([]string, error)
	UploadFile(fileType string, file *multipart.FileHeader) (string, error)
	ReadFile(directory, fileName string) ([]byte, error)
}

type fileService struct{}

func NewFileService() FileServiceImpl {
	return &fileService{}
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
