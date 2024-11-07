package category

import "github.com/labstack/echo"

func InitCategoryRoutes(e *echo.Echo, service CategoryServiceImpl) {
	h := NewCategoryHandler(service)
	r := e.Group("/categories")
	r.GET("", h.GetCategories)
	r.POST("", h.CreateCategory)

}
