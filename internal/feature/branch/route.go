package branch

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitBranchRoutes(e *echo.Echo, service BranchServiceImpl) {
	h := NewUserHandler(service)
	r := e.Group("/branches")
	r.POST("", h.CreateBranch)

	r.Use(middleware.JWTAccessMiddleware)
	r.GET("", h.GetBranches)
	r.GET("/:id", h.GetBranchByID)
	r.PATCH("/:id", h.UpdateBranch)
}
