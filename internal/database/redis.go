package database

import (
	"context"
	"inverntory_management/config"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis(cfg config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.REDIS_ADDR,
		Password: cfg.REDIS_PWD,
		DB:       cfg.REDIS_DB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return client
}
