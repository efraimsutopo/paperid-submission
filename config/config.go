package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// Server Config
	AppPort string `envconfig:"APP_PORT"`

	// Database Config
	DBHost     string `envconfig:"DB_HOST"`
	DBPort     string `envconfig:"DB_PORT"`
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`
	DBName     string `envconfig:"DB_NAME"`

	// JWT Config
	JWTSecret string `envconfig:"JWT_SECRET"`
}

var cfg Config

func New() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	if err := envconfig.Process("", &cfg); err != nil {
		return err
	}

	return nil
}

func Get() Config {
	return cfg
}
