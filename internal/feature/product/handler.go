package product

import (
	"inverntory_management/internal/exception"
	"inverntory_management/internal/types"
	custom "inverntory_management/pkg/errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ProductHandler struct {
	productService ProductServiceImpl
}

func NewProductHandler(productService ProductServiceImpl) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetProducts(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || page <= 0 {
		limit = 10
	}

	products, total, err := h.productService.FindAll(page, limit)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     products,
		Status:   http.StatusOK,
		Messages: "Success",
		Total:    &total,
	})
}

func (h *ProductHandler) GetProduct(c echo.Context) error {
	id := c.Param("id")

	product, err := h.productService.FindByID(id)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     product,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	dto := new(ProductCreateDto)
	if err := c.Bind(dto); err != nil {
		return custom.NewBadRequestError(err.Error())

	}

	if err := c.Validate(dto); err != nil {
		return custom.NewBadRequestError(err.Error())
	}

	if err := h.productService.Create(*dto); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, types.Response{
		Data:     true,
		Status:   http.StatusCreated,
		Messages: "Success",
	})
}

func (h *ProductHandler) AddVariant(c echo.Context) error {
	id := c.Param("id")

	dto := new(VariantCreateDto)
	if err := c.Bind(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := c.Validate(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := h.productService.AddVariant(id, *dto); err != nil {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Data:     err.Error(),
			Status:   http.StatusInternalServerError,
			Messages: "Error",
		})
	}

	return c.JSON(http.StatusCreated, types.Response{
		Data:     true,
		Status:   http.StatusCreated,
		Messages: "Success",
	})
}
