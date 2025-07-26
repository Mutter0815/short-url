package configs

import (
	"log"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Host     string `env:"DB_HOST"`
	Name     string `env:"DB_NAME"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

func Load() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Ошибка загрузки переменных окружения: %v", err)
	}
	return cfg
}
