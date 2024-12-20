package product

import (
	"inverntory_management/internal/exception"
	"inverntory_management/internal/types"
	err_response "inverntory_management/pkg/errors"
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
		return err
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
	dto := new(CreateProductDTO)
	if err := c.Bind(dto); err != nil {
		return err_response.NewBadRequestError(err.Error())

	}

	if err := c.Validate(dto); err != nil {
		return err_response.NewBadRequestError(err.Error())
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

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	product_id := c.Param("id")
	dto := new(UpdateProductDTO)
	if err := c.Bind(dto); err != nil {
		return err_response.NewBadRequestError(err.Error())

	}

	if err := c.Validate(dto); err != nil {
		return err_response.NewBadRequestError(err.Error())
	}

	if err := h.productService.Update(product_id, *dto); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     true,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}
