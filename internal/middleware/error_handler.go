// pkg/middleware/error_handler.go
package middleware

import (
	"errors"
	custom "inverntory_management/pkg/errors"

	"github.com/labstack/echo"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			var appErr *custom.AppError
			if !errors.As(err, &appErr) {
				appErr = custom.NewInternalServerError()
			}

			if errors.Is(err, echo.ErrNotFound) {
				appErr = custom.NewNotFoundError("404 Not found")
			}

			return c.JSON(appErr.Code, appErr)
		}
		return nil
	}
}
