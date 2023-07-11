package app

import (
	"github.com/cheeeasy2501/auth-id/pkg/database"
	grpcserver "github.com/cheeeasy2501/auth-id/pkg/grpc-server"
	"github.com/cheeeasy2501/auth-id/pkg/server"
)

type Config struct {
	Authorization IAuthorizationConfig
	Database      database.IConfig
	HTTP          server.IHTTPConfig
	GRPC          grpcserver.IGRPCConfig
}

func NewConfig() *Config {
	return &Config{
		Authorization: NewAuthorizationConfig(),
		Database:      database.NewConfig(),
		HTTP:          server.NewConfig(),
		GRPC:		   grpcserver.NewConfig(),
	}
}