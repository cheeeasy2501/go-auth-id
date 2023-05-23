package authorization

import (
	"fmt"
	"os"
)

type IConfig interface {
	GetHost() string
	GetPort() string
	GetAddr() string
}

type Config struct {
	host string
	port string
}

func NewConfig() IConfig {
	return &Config{
		host: os.Getenv("GRPC_HOST"),
		port: os.Getenv("GRPC_PORT"),
	}
}

func (cfg *Config) GetHost() string {
	return cfg.host
}

func (cfg *Config) GetPort() string {
	return cfg.port
}

func (cfg *Config) GetAddr() string {
	return fmt.Sprintf("%s:%s", cfg.GetHost(), cfg.GetPort())
}
