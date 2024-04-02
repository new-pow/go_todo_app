package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"TODO_PORT" envDefault:"80"`
}

func New() (*Config, error) {
	cfg := &Config{} // 구조체 포인터 생성
	// 환경 변수 parse
	// Parse 함수는 'env' 태그가 포함된 구조체를 구문 분석하고 환경 변수에서 해당 값을 로드합니다.
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
