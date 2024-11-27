package variant

import (
	"inverntory_management/internal/exception"
	"inverntory_management/internal/types"
	"net/http"

	"github.com/labstack/echo"
)

type VariantHandler struct {
	variantService VariantService
}

func NewProductHandler(variantService VariantService) *VariantHandler {
	return &VariantHandler{variantService: variantService}
}

func (h *VariantHandler) AddVariant(c echo.Context) error {
	product_id := c.Param("product_id")

	dto := new(VariantCreateDto)
	if err := c.Bind(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := c.Validate(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := h.variantService.Create(product_id, *dto); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, types.Response{
		Data:     true,
		Status:   http.StatusCreated,
		Messages: "Success",
	})
}

func (h *VariantHandler) Update(c echo.Context) error {
	id := c.Param("id")

	dto := new(VariantUpdateDto)
	if err := c.Bind(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := c.Validate(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := h.variantService.Update(id, *dto); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, types.Response{
		Data:     true,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}

func (h *VariantHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := h.variantService.Delete(id); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, types.Response{
		Data:     true,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}
