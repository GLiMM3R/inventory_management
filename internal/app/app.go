package app

import (
	"inverntory_management/config"
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
	user "inverntory_management/internal/feature/user"
	"inverntory_management/internal/utils"

	custom "inverntory_management/internal/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Initialize() (*echo.Echo, error) {
	// Load Configuration
	config.LoadConfig(".")

	e := echo.New()

	e.Validator = utils.NewValidator()

	// Middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(custom.ErrorHandler)

	// Initialize database
	database.InitPostgres()
	redisClient := database.InitRedis()

	// Initialize Repositories
	userRepo := user.NewUserRepository(database.DB)
	branchRepo := branch.NewBranchRepository(database.DB)
	inventoryRepo := inventory.NewInventoryRepository(database.DB)
	priceRepo := price.NewPriceRepository(database.DB)
	saleRepo := sale.NewSaleRepository(database.DB)
	transferRepo := inventory_transfer.NewInventoryTransferRepository(database.DB)
	reportRepo := report.NewReportRepository(database.DB)
	productRepo := product.NewProductRepository(database.DB)
	categoryRepo := category.NewCategoryRepository(database.DB)

	// Initialize Services
	authService := auth.NewAuthService(userRepo, redisClient)
	userService := user.NewUserService(userRepo)
	branchService := branch.NewBranchService(branchRepo, userRepo)
	inventoryService := inventory.NewInventoryService(inventoryRepo, userRepo, priceRepo)
	priceService := price.NewPriceService(priceRepo)
	saleService := sale.NewSaleService(saleRepo)
	transferService := inventory_transfer.NewInventoryService(transferRepo, userRepo)
	reportService := report.NewReportService(reportRepo, userRepo)
	productService := product.NewProductService(productRepo)
	categoryService := category.NewCategoryService(categoryRepo)
	fileService := files.NewFileService()

	// Initialize Routes
	auth.InitAuthRoutes(e, authService)
	user.InitUserRoutes(e, userService)
	branch.InitBranchRoutes(e, branchService)
	inventory.InitInventoryRoutes(e, inventoryService)
	price.InitPriceRoutes(e, priceService)
	sale.InitSaleRoutes(e, saleService)
	inventory_transfer.InitInventoryTransferRoutes(e, transferService)
	report.InitReportRoutes(e, reportService)
	product.InitProductRoutes(e, productService)
	category.InitCategoryRoutes(e, categoryService)
	files.InitFileRoutes(e, fileService)

	return e, nil
}
