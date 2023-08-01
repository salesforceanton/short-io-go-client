package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const ENV_PREFIX = "shortener"

type Config struct {
	Domain string
	Token  string
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := envconfig.Process(ENV_PREFIX, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
