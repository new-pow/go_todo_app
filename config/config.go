package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"TODO_PORT" envDefault:"80"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil { // 환경 변수 parse
		return nil, err
	}
	return cfg, nil
}
