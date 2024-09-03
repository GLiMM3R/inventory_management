package user

import (
	"inverntory_management/internal/exception"
	"inverntory_management/internal/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type UserHandler struct {
	service UserServiceImpl
}

func NewUserHandler(service UserServiceImpl) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || page <= 0 {
		limit = 10
	}

	users, total, err := h.service.GetAll(page, limit)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     users,
		Status:   http.StatusOK,
		Messages: "Success",
		Total:    &total,
	})
}

func (h *UserHandler) GetUserByUsername(c echo.Context) error {
	username := c.Param("username")
	user, err := h.service.FindByUsername(username)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     user,
		Status:   http.StatusOK,
		Messages: "Success",
	})
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	dto := new(UserCreateDto)

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
