package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	APP_NAME              string `mapstructure:"APP_NAME"`
	PORT                  int    `mapstructure:"PORT"`
	EMAIL                 string `mapstructure:"EMAIL"`
	EMAIL_PWD             string `mapstructure:"EMAIL_PWD"`
	DATABASE_DRIVER       string `mapstructure:"DATABASE_DRIVER"`
	DATABASE_URL          string `mapstructure:"DATABASE_URL"`
	ACCESS_SECRET         string `mapstructure:"ACCESS_SECRET"`
	ACCESS_EXPIRATION     int64  `mapstructure:"ACCESS_EXPIRATION"`
	REFRESH_SECRET        string `mapstructure:"REFRESH_SECRET"`
	REFRESH_EXPIRATION    int64  `mapstructure:"REFRESH_EXPIRATION"`
	RESET_SECRET          string `mapstructure:"DB_HOST"`
	RESET_EXPIRATION      int64  `mapstructure:"RESET_EXPIRATION"`
	REDIS_ADDR            string `mapstructure:"REDIS_ADDR"`
	REDIS_PWD             string `mapstructure:"DB_HOST"`
	REDIS_DB              int    `mapstructure:"REDIS_DB"`
	AWS_ACCESS_KEY_ID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWS_SECRET_ACCESS_KEY string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AWS_BUCKET_NAME       string `mapstructure:"AWS_BUCKET_NAME"`
	AWS_BUCKET_REGION     string `mapstructure:"AWS_BUCKET_REGION"`
	AWS_ENDPOINT          string `mapstructure:"AWS_ENDPOINT"`
	PRIVATE_KEY           []byte
	PUBLIC_KEY            []byte
}

var AppConfig Config

// func loadKey() {
// 	fmt.Println("Initializing configuration...")
// 	privateBytes, err := os.ReadFile("keys/private.pem")
// 	if err != nil {
// 		log.Fatalf("Error reading private key: %v", err)
// 	}

// 	AppConfig.PRIVATE_KEY = privateBytes

// 	publicBytes, err := os.ReadFile("keys/public.pem")
// 	if err != nil {
// 		log.Fatalf("Error reading public key: %v", err)
// 	}

// 	AppConfig.PUBLIC_KEY = publicBytes
// }

func LoadConfig(path, file string) Config {
	cfg := &Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName(file)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// loadKey()

	return *cfg
}
