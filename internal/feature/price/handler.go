package price

import (
	"inverntory_management/internal/exception"
	"inverntory_management/internal/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type PriceHandler struct {
	service PriceServiceImpl
}

func NewPriceHandler(service PriceServiceImpl) *PriceHandler {
	return &PriceHandler{service: service}
}

func (h *PriceHandler) GetPrices(c echo.Context) error {
	inventoryID := c.QueryParam("inventory_id")

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || page <= 0 {
		limit = 10
	}

	branches, total, err := h.service.GetAll(inventoryID, page, limit)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     branches,
		Status:   http.StatusOK,
		Messages: "Success",
		Total:    &total,
	})
}

func (h *PriceHandler) GetPriceByID(c echo.Context) error {
	id := c.Param("id")
	user, err := h.service.FindByID(id)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     user,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}

func (h *PriceHandler) CreatePrice(c echo.Context) error {
	dto := new(PriceCreateDto)

	if err := c.Bind(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := c.Validate(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := h.service.Create(*dto); err != nil {
		return exception.HandleError(c, err)
	}

	return c.JSON(http.StatusCreated, types.Response{
		Data:     true,
		Status:   http.StatusCreated,
		Messages: "Success",
	})
}

func (h *PriceHandler) UpdatePrice(c echo.Context) error {
	inventory_id := c.Param("id")

	dto := new(PriceUpdateDto)

	if err := c.Bind(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := c.Validate(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := h.service.Update(inventory_id, *dto); err != nil {
		return exception.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     true,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}
