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
	default:
		return c.JSON(http.StatusInternalServerError, types.Response{
			Data:     nil,
			Status:   http.StatusInternalServerError,
			Messages: ErrInternal.Error(),
		})
	}
}
