package files

import (
	"inverntory_management/internal/types"
	custom "inverntory_management/pkg/errors"
	"net/http"

	"github.com/labstack/echo"
)

type FileHandler struct {
	service FileServiceImpl
}

func NewFileHandler(service FileServiceImpl) FileHandler {
	return FileHandler{service: service}
}

func (h *FileHandler) UploadFile(c echo.Context) error {
	fileType := c.QueryParam("type")

	file, err := c.FormFile("file")
	if err != nil {
		return custom.NewBadRequestError(err.Error())
	}

	fileName, err := h.service.UploadFile(fileType, file)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     fileName,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}

func (h *FileHandler) GetFile(c echo.Context) error {
	directory := c.Param("directory")
	fileName := c.Param("name")
	file, err := h.service.ReadFile(directory, fileName)
	if err != nil {
		return err
	}

	// Assuming the file is a path to the image file on the server
	return c.Blob(http.StatusOK, "image/jpeg", file)
}
