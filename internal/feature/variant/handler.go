package variant

import (
	"inverntory_management/internal/types"
	err_response "inverntory_management/pkg/errors"
	"net/http"

	"github.com/labstack/echo"
)

type VariantHandler struct {
	variantService VariantService
}

func NewVariantHandler(variantService VariantService) *VariantHandler {
	return &VariantHandler{variantService: variantService}
}

func (h *VariantHandler) GetVariant(c echo.Context) error {
	id := c.Param("id")

	variant, err := h.variantService.FindByID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     variant,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}

func (h *VariantHandler) CreateVariant(c echo.Context) error {
	product_id := c.Param("product_id")

	dto := new(CreateVariantDTO)
	if err := c.Bind(dto); err != nil {
		return err_response.NewBadRequestError(err.Error())
	}

	if err := c.Validate(dto); err != nil {
		return err_response.NewBadRequestError(err.Error())
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

func (h *VariantHandler) UpdateVariant(c echo.Context) error {
	id := c.Param("id")

	dto := new(UpdateVariantDTO)
	if err := c.Bind(dto); err != nil {
		return err_response.NewBadRequestError(err.Error())
	}

	if err := c.Validate(dto); err != nil {
		return err_response.NewBadRequestError(err.Error())
	}

	if err := h.variantService.Update(id, *dto); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     true,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}

func (h *VariantHandler) DeleteVariant(c echo.Context) error {
	id := c.Param("id")

	if err := h.variantService.Delete(id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     true,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}
