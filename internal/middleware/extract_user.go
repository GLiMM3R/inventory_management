package middleware

import (
	"inverntory_management/internal/exception"
	"inverntory_management/internal/types"

	"github.com/labstack/echo"
)

func ExtractUser(c echo.Context) (*types.UserClaims, error) {
	// Get the user from the context
	userClaims, ok := c.Get("user").(*types.UserClaims)
	if !ok {
		return nil, exception.ErrUnauthorized
	}

	return userClaims, nil
}
