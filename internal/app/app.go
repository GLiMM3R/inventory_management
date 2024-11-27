package app

import (
	"inverntory_management/config"
	"inverntory_management/internal/database"
	"inverntory_management/internal/utils"
	aws_service "inverntory_management/pkg/aws"
	"log"

	custom "inverntory_management/internal/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Initialize() (*echo.Echo, error) {
	// Load Configuration
	config.LoadConfig(".", ".env")

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

	s3Client, err := aws_service.NewS3Client(aws_service.Config{
		Region:          config.AppConfig.AWS_BUCKET_REGION,
		Endpoint:        config.AppConfig.AWS_ENDPOINT,
		AccessKeyID:     config.AppConfig.AWS_ACCESS_KEY_ID,
		SecretAccessKey: config.AppConfig.AWS_SECRET_ACCESS_KEY})
	if err != nil {
		log.Println("Error initializing S3 client: ", err)
		return nil, err
	}

	// Initialize Repositories and Services
	repo := initializeRepositories(redisClient)
	services := initializeServices(repo, s3Client)

	// Initialize Routes
	initializeRoutes(e, services)

	return e, nil
}
