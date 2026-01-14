package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	APP_PORT       string `env:"APP_PORT"`
	DB_HOST        string `env:"DB_HOST"`
	DB_PORT        string `env:"DB_PORT"`
	DB_DATABASE    string `env:"DB_DATABASE"`
	DB_USERNAME    string `env:"DB_USERNAME"`
	DB_PASSWORD    string `env:"DB_PASSWORD"`
	JWT_SECRET     string `env:"JWT_SECRET"`
	JWT_EXPIRES    int    `env:"JWT_EXPIRES"`
	GEMINI_API_KEY string `env:"GEMINI_API_KEY"`
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := new(Config)
	err = env.Parse(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
