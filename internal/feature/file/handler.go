package files

import (
	"inverntory_management/internal/types"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type FileHandler struct {
	// service FileServiceImpl
}

func NewFileHandler() FileHandler {
	return FileHandler{}
}

func (h *FileHandler) UploadFile(c echo.Context) error {
	fileType := c.QueryParam("type")

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Data:     nil,
			Status:   http.StatusNotFound,
			Messages: err.Error(),
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Data:     nil,
			Status:   http.StatusNotFound,
			Messages: err.Error(),
		})
	}
	defer src.Close()

	fileName := uuid.NewString() + filepath.Ext(file.Filename)

	dstPath := filepath.Join("uploads", fileType, fileName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Data:     nil,
			Status:   http.StatusNotFound,
			Messages: err.Error(),
		})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Data:     nil,
			Status:   http.StatusNotFound,
			Messages: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     fileName,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}
