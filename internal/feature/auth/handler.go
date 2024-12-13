package auth

import (
	"inverntory_management/internal/exception"
	"inverntory_management/internal/middleware"
	"inverntory_management/internal/types"
	"inverntory_management/pkg/errors"
	"net/http"

	"github.com/labstack/echo"
)

type AuthHandler struct {
	service AuthServiceImpl
}

func NewAuthHandler(service AuthServiceImpl) AuthHandler {
	return AuthHandler{service: service}
}

func (h *AuthHandler) Login(c echo.Context) error {
	request := new(AuthRequest)
	if err := c.Bind(request); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	if err := c.Validate(request); err != nil {
		return errors.NewInternalServerError()
	}

	response, err := h.service.Login(request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     response,
		Status:   http.StatusOK,
		Messages: "Logged in successfully",
	})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return exception.HandleError(c, exception.ErrUnauthorized)
	}

	if err := h.service.Logout(token); err != nil {
		return exception.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     true,
		Status:   http.StatusOK,
		Messages: "Logged out successfully",
	})
}

func (h *AuthHandler) GetRefreshToken(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return exception.HandleError(c, exception.ErrUnauthorized)
	}

	userClaims, err := middleware.ExtractUser(c)
	if err != nil {
		return exception.HandleError(c, err)
	}

	response, err := h.service.GetRefreshToken(token, userClaims)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     response,
		Status:   http.StatusOK,
		Messages: "Successfully",
	})
}

func (h *AuthHandler) SendOTP(c echo.Context) error {
	request := new(OTPRequest)
	if err := c.Bind(request); err != nil {
		return exception.HandleError(c, exception.ErrInvalidData)
	}

	if err := h.service.SendOTP(request.Username); err != nil {
		return exception.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, types.Response{
		Data:     true,
		Status:   http.StatusOK,
		Messages: "Successfully",
	})
}

func (h *AuthHandler) CheckAuth(c echo.Context) error {
	return c.JSON(http.StatusOK, types.Response{
		Data:     true,
		Status:   http.StatusOK,
		Messages: "Successfully",
	})
}
