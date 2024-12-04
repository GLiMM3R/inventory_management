package app

import (
	"inverntory_management/internal/feature/auth"
	"inverntory_management/internal/feature/category"
	files "inverntory_management/internal/feature/file"
	"inverntory_management/internal/feature/media"
	"inverntory_management/internal/feature/product"
	"inverntory_management/internal/feature/report"
	"inverntory_management/internal/feature/sale"
	"inverntory_management/internal/feature/user"
	"inverntory_management/internal/feature/variant"
)

type AppServices struct {
	AuthService     auth.AuthServiceImpl
	UserService     user.UserServiceImpl
	SaleService     sale.SaleServiceImpl
	ReportService   report.ReportServiceImpl
	ProductService  product.ProductServiceImpl
	CategoryService category.CategoryServiceImpl
	VariantService  variant.VariantService
	FileService     files.FileServiceImpl
	MediaService    media.MediaService
}

// initializeServices initializes all services and returns them in a struct.
func (a *App) initServices(repo *AppRepositories) *AppServices {
	return &AppServices{
		AuthService:     auth.NewAuthService(repo.UserRepo, &repo.RedisClient),
		UserService:     user.NewUserService(repo.UserRepo),
		SaleService:     sale.NewSaleService(repo.SaleRepo),
		ReportService:   report.NewReportService(repo.ReportRepo, repo.UserRepo),
		ProductService:  product.NewProductService(repo.ProductRepo, repo.MediaRepo, *a.s3Client, a.cfg),
		CategoryService: category.NewCategoryService(repo.CategoryRepo),
		MediaService:    media.NewMediaService(repo.MediaRepo, *a.s3Client),
		VariantService:  variant.NewVariantService(repo.VariantRepo, repo.ProductRepo, *a.s3Client, a.cfg),
		FileService:     files.NewFileService(*a.s3Client, a.cfg),
	}
}
