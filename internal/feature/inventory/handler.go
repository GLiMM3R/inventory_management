package inventory

import (
	"inverntory_management/internal/exception"
	"inverntory_management/internal/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type InventoryHandler struct {
	service InventoryServiceImpl
}

func NewInventoryHandler(service InventoryServiceImpl) *InventoryHandler {
	return &InventoryHandler{service: service}
}

func (h *InventoryHandler) GetInventories(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || page <= 0 {
		limit = 10
	}

	branches, total, err := h.service.GetAll(page, limit)
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

func (h *InventoryHandler) GetInventoryByID(c echo.Context) error {
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

func (h *InventoryHandler) CreateInventory(c echo.Context) error {
	dto := new(InventoryCreateDto)

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

func (h *InventoryHandler) UpdateInventory(c echo.Context) error {
	inventory_id := c.Param("id")

	dto := new(InventoryUpdateDto)

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
