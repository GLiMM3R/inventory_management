package app

import (
	"inverntory_management/internal/database"
	"inverntory_management/internal/feature/auth"
	"inverntory_management/internal/feature/category"
	files "inverntory_management/internal/feature/file"
	"inverntory_management/internal/feature/media"
	"inverntory_management/internal/feature/product"
	"inverntory_management/internal/feature/report"
	"inverntory_management/internal/feature/sale"
	"inverntory_management/internal/feature/user"
	aws_service "inverntory_management/pkg/aws"

	"github.com/labstack/echo"
	"github.com/redis/go-redis/v9"
)

type AppRepositories struct {
	UserRepo     user.UserRepositoryImpl
	SaleRepo     sale.SaleRepositoryImpl
	ReportRepo   report.ReportRepositoryImpl
	ProductRepo  product.ProductRepositoryImpl
	CategoryRepo category.CategoryRepositoryImpl
	MediaRepo    media.MediaRepository
	RedisClient  redis.Client
}

type AppServices struct {
	AuthService     auth.AuthServiceImpl
	UserService     user.UserServiceImpl
	SaleService     sale.SaleServiceImpl
	ReportService   report.ReportServiceImpl
	ProductService  product.ProductServiceImpl
	CategoryService category.CategoryServiceImpl
	FileService     files.FileServiceImpl
	MediaService    media.MediaService
}

// initializeRepositories initializes all repositories and returns dependencies struct.
func initializeRepositories(redisClient *redis.Client) *AppRepositories {
	return &AppRepositories{
		UserRepo:     user.NewUserRepository(database.DB),
		SaleRepo:     sale.NewSaleRepository(database.DB),
		ReportRepo:   report.NewReportRepository(database.DB),
		ProductRepo:  product.NewProductRepository(database.DB),
		CategoryRepo: category.NewCategoryRepository(database.DB),
		MediaRepo:    media.NewMediaRepository(database.DB),
		RedisClient:  *redisClient,
	}
}

// initializeServices initializes all services and returns them in a struct.
func initializeServices(repo *AppRepositories, s3Client *aws_service.S3Client) *AppServices {
	return &AppServices{
		AuthService:     auth.NewAuthService(repo.UserRepo, &repo.RedisClient),
		UserService:     user.NewUserService(repo.UserRepo),
		SaleService:     sale.NewSaleService(repo.SaleRepo),
		ReportService:   report.NewReportService(repo.ReportRepo, repo.UserRepo),
		ProductService:  product.NewProductService(repo.ProductRepo, repo.MediaRepo, *s3Client),
		CategoryService: category.NewCategoryService(repo.CategoryRepo),
		MediaService:    media.NewMediaService(repo.MediaRepo, *s3Client),
		FileService:     files.NewFileService(*s3Client),
	}
}

// initializeRoutes initializes routes for each service.
func initializeRoutes(e *echo.Echo, service *AppServices) {
	auth.InitAuthRoutes(e, service.AuthService)
	user.InitUserRoutes(e, service.UserService)
	sale.InitSaleRoutes(e, service.SaleService)
	report.InitReportRoutes(e, service.ReportService)
	product.InitProductRoutes(e, service.ProductService)
	category.InitCategoryRoutes(e, service.CategoryService)
	files.InitFileRoutes(e, service.FileService)
}
