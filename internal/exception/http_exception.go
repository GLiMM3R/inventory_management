package exception

import (
	"inverntory_management/internal/types"
	"net/http"

	"github.com/labstack/echo"
)

func HandleError(c echo.Context, err error) error {
	switch err {
	case ErrNotFound:
		return c.JSON(http.StatusNotFound, types.Response{
			Data:     nil,
			Status:   http.StatusNotFound,
			Messages: ErrNotFound.Error(),
		})
	case ErrDuplicateEntry:
		return c.JSON(http.StatusBadRequest, types.Response{
			Data:     nil,
			Status:   http.StatusBadRequest,
			Messages: ErrDuplicateEntry.Error(),
		})
	case ErrInsufficientQuantity:
		return c.JSON(http.StatusBadRequest, types.Response{
			Data:     nil,
			Status:   http.StatusBadRequest,
			Messages: ErrInsufficientQuantity.Error(),
		})
	case ErrInvalidToken:
		return c.JSON(http.StatusUnauthorized, types.Response{
			Data:     nil,
			Status:   http.StatusUnauthorized,
			Messages: ErrInvalidToken.Error(),
		})
	case ErrTokenNotFound:
		return c.JSON(http.StatusUnauthorized, types.Response{
			Data:     nil,
			Status:   http.StatusUnauthorized,
			Messages: ErrTokenNotFound.Error(),
		})
	case ErrInvalidCredentials:
		return c.JSON(http.StatusUnauthorized, types.Response{
			Data:     nil,
			Status:   http.StatusUnauthorized,
			Messages: ErrInvalidCredentials.Error(),
		})
	case ErrTokenExpired:
		return c.JSON(http.StatusUnauthorized, types.Response{
			Data:     nil,
			Status:   http.StatusUnauthorized,
			Messages: ErrTokenExpired.Error(),
		})
	default:
		return c.JSON(http.StatusInternalServerError, types.Response{
			Data:     nil,
			Status:   http.StatusInternalServerError,
			Messages: ErrInternal.Error(),
		})
	}
}
