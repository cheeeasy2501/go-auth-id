package app

import (
	"github.com/cheeeasy2501/auth-id/pkg/database"
	"github.com/cheeeasy2501/auth-id/pkg/server"
)

type Config struct {
	Authorization IAuthorizationConfig
	Database      database.IConfig
	HTTP      	  server.IHTTPConfig	
}

func NewConfig() *Config {
	return &Config{
		Authorization: NewAuthorizationConfig(),
		Database: database.NewPostgresConfig(),
		HTTP: server.NewHTTPConfig(),
	}
}