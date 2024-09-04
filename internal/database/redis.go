package database

import (
	"context"
	"inverntory_management/config"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.REDIS_ADDR,
		Password: config.AppConfig.REDIS_PWD,
		DB:       config.AppConfig.REDIS_DB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return client
}
