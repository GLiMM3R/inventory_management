package app

import (
	"inverntory_management/internal/feature/category"
	"inverntory_management/internal/feature/media"
	"inverntory_management/internal/feature/product"
	"inverntory_management/internal/feature/report"
	"inverntory_management/internal/feature/sale"
	"inverntory_management/internal/feature/user"
	"inverntory_management/internal/feature/variant"

	"github.com/redis/go-redis/v9"
)

type AppRepositories struct {
	UserRepo     user.UserRepositoryImpl
	SaleRepo     sale.SaleRepositoryImpl
	ReportRepo   report.ReportRepositoryImpl
	ProductRepo  product.ProductRepositoryImpl
	CategoryRepo category.CategoryRepositoryImpl
	MediaRepo    media.MediaRepository
	VariantRepo  variant.VariantRepository
	RedisClient  redis.Client
}

// initializeRepositories initializes all repositories and returns dependencies struct.
func (a *App) initRepositories() *AppRepositories {
	return &AppRepositories{
		UserRepo:     user.NewUserRepository(a.db),
		SaleRepo:     sale.NewSaleRepository(a.db),
		ReportRepo:   report.NewReportRepository(a.db),
		ProductRepo:  product.NewProductRepository(a.db),
		CategoryRepo: category.NewCategoryRepository(a.db),
		MediaRepo:    media.NewMediaRepository(a.db),
		VariantRepo:  variant.NewVariantRepository(a.db),
		RedisClient:  *a.redisClient,
	}
}
