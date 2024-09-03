package app

import (
	"inverntory_management/config"
	"inverntory_management/internal/database"
	"inverntory_management/internal/feature/branch"
	"inverntory_management/internal/feature/inventory"
	user "inverntory_management/internal/feature/user"
	"inverntory_management/internal/utils"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Initialize() (*echo.Echo, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}
	log.Printf("Working directory: %s", dir)
	// Load Configuration
	config.LoadConfig(".")

	e := echo.New()

	e.Validator = utils.NewValidator()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize database
	database.InitPostgres()

	// Initialize Repositories
	userRepo := user.NewUserRepository(database.DB)
	branchRepo := branch.NewBranchRepository(database.DB)
	inventoryRepo := inventory.NewInventoryRepository(database.DB)

	// Initialize Services
	userService := user.NewUserService(userRepo)
	branchService := branch.NewBranchService(branchRepo)
	inventoryService := inventory.NewInventoryService(inventoryRepo)

	// Initialize Routes
	user.InitUserRoutes(e, userService)
	branch.InitBranchRoutes(e, branchService)
	inventory.InitInventoryRoutes(e, inventoryService)

	return e, nil
}
