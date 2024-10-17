package env

import (
	"errors"
	"github.com/vakhrushevk/auth-service/internal/config"
	"net"
	"os"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

const (
	grpcHostEnvNome = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type grpcConfig struct {
	host string
	port string
}

// Address Метод Address() объединяет значения хоста и порта из конфигурации grpcConfig в строку формата "host
// " с использованием net.JoinHostPort. Возвращает полный адрес для подключения к gRPC сервису.
func (g *grpcConfig) Address() string {
	return net.JoinHostPort(g.host, g.port)
}

// NewGRPCConfig Функция NewGRPCConfig создаёт и возвращает конфигурацию для gRPC, извлекая значения хоста и порта из переменных окружения. Если хост или порт не найдены, возвращает соответствующую ошибку.
func NewGRPCConfig() (config.GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvNome)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}
