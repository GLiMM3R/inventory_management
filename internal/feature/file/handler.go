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

func (h *FileHandler) UploadMultiFiles(c echo.Context) error {
	fileType := c.QueryParam("type")

	form, err := c.MultipartForm()
	if err != nil {
		return custom.NewBadRequestError(err.Error())
	}

	files := form.File["files"]
	if len(files) == 0 {
		return custom.NewBadRequestError("No files provided")
	}

	fileNames, err := h.service.UploadMultiFiles(fileType, files)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     fileNames,
		Status:   http.StatusOK,
		Messages: "Success",
	})
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

func (h *FileHandler) GeneratePresignPutObject(c echo.Context) error {
	var request PutObjectRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	res, err := h.service.GeneratePresignPutObject(request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     res,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}
