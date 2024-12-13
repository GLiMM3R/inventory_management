package app

import (
	"context"
	"fmt"
	"inverntory_management/config"
	"inverntory_management/internal/database"
	custom_middlware "inverntory_management/internal/middleware"
	"inverntory_management/internal/utils"
	aws_service "inverntory_management/pkg/aws"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	echo        *echo.Echo
	cfg         config.Config
	db          *gorm.DB
	redisClient *redis.Client
	s3Client    *aws_service.S3Client
}

func New(cfg config.Config) *App {
	s3Client, err := aws_service.NewS3Client(aws_service.Config{
		Region:          cfg.AWS_BUCKET_REGION,
		Endpoint:        cfg.AWS_ENDPOINT,
		AccessKeyID:     cfg.AWS_ACCESS_KEY_ID,
		SecretAccessKey: cfg.AWS_SECRET_ACCESS_KEY,
	})

	if err != nil {
		log.Println("Error initializing S3 client: ", err)
	}

	app := &App{
		echo:        echo.New(),
		cfg:         cfg,
		db:          database.InitPostgres(cfg),
		redisClient: database.InitRedis(cfg),
		s3Client:    s3Client,
	}

	return app
}

func (a *App) Start(ctx context.Context) {
	port := fmt.Sprintf("0.0.0.0:%d", a.cfg.PORT)

	a.echo.Validator = utils.NewValidator()

	// Middleware
	a.echo.Use(middleware.CORS())
	a.echo.Use(middleware.Logger())
	a.echo.Use(middleware.Recover())
	a.echo.Use(custom_middlware.ErrorHandler)

	err := a.redisClient.Ping(ctx).Err()
	if err != nil {
		log.Println("Error pinging Redis: ", err)
	}

	repo := a.initRepositories()
	services := a.initServices(repo)
	a.initRoutes(services)

	fmt.Println("Starting server...")

	// Start server
	go func() {
		if err := a.echo.Start(port); err != nil && err != http.ErrServerClosed {
			a.echo.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.echo.Shutdown(ctx); err != nil {
		a.echo.Logger.Fatal(err)
	}
}
