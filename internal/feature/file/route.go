package files

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitFileRoutes(e *echo.Echo, service FileServiceImpl) {
	h := NewFileHandler(service)
	r := e.Group("/files")
	r.Use(middleware.JWTAccessMiddleware)
	r.POST("", h.UploadFile)
	r.GET("/:directory/:name", h.GetFile)
}
