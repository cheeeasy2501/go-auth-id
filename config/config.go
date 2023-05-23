package app

import (
	"github.com/cheeeasy2501/auth-id/pkg/database"
	auth_pkg "github.com/cheeeasy2501/auth-id/pkg/grpc/v1/authorization"
	"github.com/cheeeasy2501/auth-id/pkg/server"
)

type Config struct {
	Authorization IAuthorizationConfig
	Database      database.IConfig
	HTTP          server.IHTTPConfig
	GRPC          auth_pkg.IConfig
}

func NewConfig() *Config {
	return &Config{
		Authorization: NewAuthorizationConfig(),
		Database:      database.NewPostgresConfig(),
		HTTP:          server.NewHTTPConfig(),
		GRPC:          auth_pkg.NewConfig(),
	}
}
