package branch

import "github.com/labstack/echo"

func InitBranchRoutes(e *echo.Echo, service BranchServiceImpl) {
	h := NewUserHandler(service)
	r := e.Group("/branches")
	r.GET("", h.GetBranches)
	r.GET("/:id", h.GetBranchByID)
	r.POST("", h.CreateBranch)
	r.PATCH("/:id", h.UpdateBranch)
}
