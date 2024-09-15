package inventory_transfer

import (
	"inverntory_management/internal/exception"
	"inverntory_management/internal/middleware"
	"inverntory_management/internal/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type InventoryTransactionHandler struct {
	service InventoryTransferServiceImpl
}

func NewInventoryTransactionHandler(service InventoryTransferServiceImpl) *InventoryTransactionHandler {
	return &InventoryTransactionHandler{service: service}
}

func (h *InventoryTransactionHandler) GetInventoryTransfers(c echo.Context) error {
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

func (h *InventoryTransactionHandler) GetInventoryTransferByID(c echo.Context) error {
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

func (h *InventoryTransactionHandler) CreateInventoryTransfer(c echo.Context) error {
	userClaims, err := middleware.ExtractUser(c)
	if err != nil {
		return exception.HandleError(c, err)
	}

	dto := new(InventoryTransferCreateDto)

	if err := c.Bind(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := c.Validate(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := h.service.Create(*dto, *userClaims); err != nil {
		// return exception.HandleError(c, err)
		return c.JSON(http.StatusInternalServerError, types.Response{
			Data:     true,
			Status:   http.StatusCreated,
			Messages: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, types.Response{
		Data:     true,
		Status:   http.StatusCreated,
		Messages: "Success",
	})
}
