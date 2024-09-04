package middleware

import (
	"inverntory_management/internal/types"
	"net/http"

	"github.com/labstack/echo"
)

func ExtractUser(c echo.Context) (*types.UserClaims, error) {
	// Get the user from the context
	userClaims, ok := c.Get("user").(*types.UserClaims)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
	}

	return userClaims, nil
}
