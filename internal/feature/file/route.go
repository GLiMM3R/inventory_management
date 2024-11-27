package files

import (
	"github.com/labstack/echo"
)

func InitFileRoutes(e *echo.Echo, service FileServiceImpl) {
	h := NewFileHandler(service)
	r := e.Group("/files")
	// r.Use(middleware.JWTAccessMiddleware)
	r.POST("", h.UploadFile)
	r.POST("/multiple", h.UploadMultiFiles)
	r.GET("/:directory/:name", h.GetFile)
	r.POST("/generate-presign", h.GeneratePresignPutObject)
}
