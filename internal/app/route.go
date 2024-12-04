package app

import (
	"inverntory_management/internal/feature/auth"
	"inverntory_management/internal/feature/category"
	files "inverntory_management/internal/feature/file"
	"inverntory_management/internal/feature/product"
	"inverntory_management/internal/feature/report"
	"inverntory_management/internal/feature/sale"
	"inverntory_management/internal/feature/user"
	"inverntory_management/internal/feature/variant"
)

// initializeRoutes initializes routes for each service.
func (a *App) initRoutes(service *AppServices) {
	auth.InitAuthRoutes(a.echo, service.AuthService)
	user.InitUserRoutes(a.echo, service.UserService)
	sale.InitSaleRoutes(a.echo, service.SaleService)
	report.InitReportRoutes(a.echo, service.ReportService)
	product.InitProductRoutes(a.echo, service.ProductService)
	category.InitCategoryRoutes(a.echo, service.CategoryService)
	variant.InitVariantRoutes(a.echo, service.VariantService)
	files.InitFileRoutes(a.echo, service.FileService)
}
