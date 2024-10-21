package files

import (
	"inverntory_management/internal/middleware"

	"github.com/labstack/echo"
)

func InitFileRoutes(e *echo.Echo) {
	h := NewFileHandler()
	r := e.Group("/files")
	r.Use(middleware.JWTAccessMiddleware)
	r.POST("", h.UploadFile)
}
