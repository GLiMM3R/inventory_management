package product

import (
	"inverntory_management/internal/exception"
	"inverntory_management/internal/types"
	"net/http"

	"github.com/labstack/echo"
)

type ProductHandler struct {
	productService ProductServiceImpl
}

func NewProductHandler(productService ProductServiceImpl) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	dto := new(ProductCreateDto)
	if err := c.Bind(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := c.Validate(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := h.productService.Create(*dto); err != nil {
		return exception.HandleError(c, err)
	}

	return c.JSON(http.StatusCreated, types.Response{
		Data:     true,
		Status:   http.StatusCreated,
		Messages: "Success",
	})
}
