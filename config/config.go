package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Domain string
	Token  string
}

func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("shortener", &cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
