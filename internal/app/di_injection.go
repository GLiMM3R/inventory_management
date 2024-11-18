package app

import (
	"inverntory_management/internal/database"
	"inverntory_management/internal/feature/auth"
	"inverntory_management/internal/feature/branch"
	"inverntory_management/internal/feature/category"
	files "inverntory_management/internal/feature/file"
	"inverntory_management/internal/feature/inventory"
	"inverntory_management/internal/feature/inventory_transfer"
	"inverntory_management/internal/feature/price"
	"inverntory_management/internal/feature/product"
	"inverntory_management/internal/feature/report"
	"inverntory_management/internal/feature/sale"
	"inverntory_management/internal/feature/user"

	"github.com/labstack/echo"
	"github.com/redis/go-redis/v9"
)

type AppRepositories struct {
	UserRepo      user.UserRepositoryImpl
	BranchRepo    branch.BranchRepositoryImpl
	InventoryRepo inventory.InventoryRepositoryImpl
	PriceRepo     price.PriceRepositoryImpl
	SaleRepo      sale.SaleRepositoryImpl
	TransferRepo  inventory_transfer.InventoryTransferRepositoryImpl
	ReportRepo    report.ReportRepositoryImpl
	ProductRepo   product.ProductRepositoryImpl
	CategoryRepo  category.CategoryRepositoryImpl
	RedisClient   redis.Client
}

type AppServices struct {
	AuthService      auth.AuthServiceImpl
	UserService      user.UserServiceImpl
	BranchService    branch.BranchServiceImpl
	InventoryService inventory.InventoryServiceImpl
	PriceService     price.PriceServiceImpl
	SaleService      sale.SaleServiceImpl
	TransferService  inventory_transfer.InventoryTransferServiceImpl
	ReportService    report.ReportServiceImpl
	ProductService   product.ProductServiceImpl
	CategoryService  category.CategoryServiceImpl
	FileService      files.FileServiceImpl
}

// initializeRepositories initializes all repositories and returns dependencies struct.
func initializeRepositories(redisClient *redis.Client) *AppRepositories {
	return &AppRepositories{
		UserRepo:      user.NewUserRepository(database.DB),
		BranchRepo:    branch.NewBranchRepository(database.DB),
		InventoryRepo: inventory.NewInventoryRepository(database.DB),
		PriceRepo:     price.NewPriceRepository(database.DB),
		SaleRepo:      sale.NewSaleRepository(database.DB),
		TransferRepo:  inventory_transfer.NewInventoryTransferRepository(database.DB),
		ReportRepo:    report.NewReportRepository(database.DB),
		ProductRepo:   product.NewProductRepository(database.DB),
		CategoryRepo:  category.NewCategoryRepository(database.DB),
		RedisClient:   *redisClient,
	}
}

// initializeServices initializes all services and returns them in a struct.
func initializeServices(repo *AppRepositories) *AppServices {
	return &AppServices{
		AuthService:      auth.NewAuthService(repo.UserRepo, &repo.RedisClient),
		UserService:      user.NewUserService(repo.UserRepo),
		BranchService:    branch.NewBranchService(repo.BranchRepo, repo.UserRepo),
		InventoryService: inventory.NewInventoryService(repo.InventoryRepo, repo.UserRepo, repo.PriceRepo),
		PriceService:     price.NewPriceService(repo.PriceRepo),
		SaleService:      sale.NewSaleService(repo.SaleRepo),
		TransferService:  inventory_transfer.NewInventoryService(repo.TransferRepo, repo.UserRepo),
		ReportService:    report.NewReportService(repo.ReportRepo, repo.UserRepo),
		ProductService:   product.NewProductService(repo.ProductRepo),
		CategoryService:  category.NewCategoryService(repo.CategoryRepo),
		FileService:      files.NewFileService(),
	}
}

// initializeRoutes initializes routes for each service.
func initializeRoutes(e *echo.Echo, service *AppServices) {
	auth.InitAuthRoutes(e, service.AuthService)
	user.InitUserRoutes(e, service.UserService)
	branch.InitBranchRoutes(e, service.BranchService)
	inventory.InitInventoryRoutes(e, service.InventoryService)
	price.InitPriceRoutes(e, service.PriceService)
	sale.InitSaleRoutes(e, service.SaleService)
	inventory_transfer.InitInventoryTransferRoutes(e, service.TransferService)
	report.InitReportRoutes(e, service.ReportService)
	product.InitProductRoutes(e, service.ProductService)
	category.InitCategoryRoutes(e, service.CategoryService)
	files.InitFileRoutes(e, service.FileService)
}
