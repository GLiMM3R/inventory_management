package branch

import (
	"inverntory_management/internal/exception"
	"inverntory_management/internal/middleware"
	"inverntory_management/internal/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type BranchHandler struct {
	service BranchServiceImpl
}

func NewUserHandler(service BranchServiceImpl) *BranchHandler {
	return &BranchHandler{service: service}
}

func (h *BranchHandler) GetBranches(c echo.Context) error {
	userClaims, err := middleware.ExtractUser(c)
	if err != nil {
		return exception.HandleError(c, err)
	}

	notSelf := c.QueryParam("not_self") == "true"

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || page <= 0 {
		limit = 10
	}

	branches, total, err := h.service.GetAll(page, limit, *userClaims, notSelf)
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

func (h *BranchHandler) GetBranchByID(c echo.Context) error {
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

func (h *BranchHandler) CreateBranch(c echo.Context) error {
	dto := new(BranchCreateDto)

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

func (h *BranchHandler) UpdateBranch(c echo.Context) error {
	branch_id := c.Param("id")

	dto := new(BranchUpdateDto)

	if err := c.Bind(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := c.Validate(dto); err != nil {
		return exception.HandleError(c, err)
	}

	if err := h.service.Update(branch_id, *dto); err != nil {
		return exception.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     true,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}
