package category

import (
	"inverntory_management/internal/types"
	custom "inverntory_management/pkg/errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type CategoryHandler struct {
	categoryService CategoryServiceImpl
}

func NewCategoryHandler(categoryService CategoryServiceImpl) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) GetCategories(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || page <= 0 {
		limit = 10
	}

	categories, total, err := h.categoryService.GetAll(page, limit)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     categories,
		Status:   http.StatusOK,
		Messages: "Success",
		Total:    &total,
	})
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	dto := new(CategoryRequest)
	if err := c.Bind(dto); err != nil {
		return custom.NewBadRequestError(err.Error())
	}

	if err := c.Validate(dto); err != nil {
		return custom.NewBadRequestError(err.Error())
	}

	if err := h.categoryService.Create(*dto); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, types.Response{
		Data:     nil,
		Status:   http.StatusCreated,
		Messages: "Category created successfully",
		Total:    nil,
	})
}
