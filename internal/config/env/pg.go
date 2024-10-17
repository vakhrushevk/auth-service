package env

import (
	"errors"
	"github.com/vakhrushevk/auth-service/internal/config"
	"os"
)

var _ config.PgConfig = (*pgConfig)(nil)

const (
	dnsEnvName = "PG_DSN"
)

type pgConfig struct {
	dsn string
}

// NewPGConfig Функция NewPGConfig предназначена для создания и инициализации конфигурации для подключения к базе данных PostgreSQL на основе переменной окружения.
func NewPGConfig() (config.PgConfig, error) {
	dsn := os.Getenv(dnsEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
