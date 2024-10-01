package branch

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitBranchRoutes(e *echo.Echo, service BranchServiceImpl) {
	h := NewUserHandler(service)
	r := e.Group("/branches")
	r.POST("", h.CreateBranch)

	protected := r.Group("")
	protected.Use(middleware.JWTAccessMiddleware)
	protected.GET("", h.GetBranches)
	protected.GET("/:id", h.GetBranchByID)
	protected.PATCH("/:id", h.UpdateBranch)
}
