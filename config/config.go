package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	APP_NAME           string `mapstructure:"APP_NAME"`
	PORT               int    `mapstructure:"PORT"`
	DATABASE_DRIVER    string `mapstructure:"DATABASE_DRIVER"`
	DATABASE_URL       string `mapstructure:"DATABASE_URL"`
	ACCESS_SECRET      string `mapstructure:"ACCESS_SECRET"`
	ACCESS_EXPIRATION  int64  `mapstructure:"ACCESS_EXPIRATION"`
	REFRESH_SECRET     string `mapstructure:"REFRESH_SECRET"`
	REFRESH_EXPIRATION int64  `mapstructure:"REFRESH_EXPIRATION"`
	RESET_SECRET       string `mapstructure:"DB_HOST"`
	RESET_EXPIRATION   int64  `mapstructure:"RESET_EXPIRATION"`
	REDIS_ADDR         string `mapstructure:"REDIS_ADDR"`
	REDIS_PWD          string `mapstructure:"DB_HOST"`
	REDIS_DB           int    `mapstructure:"REDIS_DB"`
}

var AppConfig Config

func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}