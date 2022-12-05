package app

import "github.com/cheeeasy2501/auth-id/package/database"

type Config struct {
	Authorization IAuthorizationConfig
	Database      database.IConfig
}

func NewConfig() *Config {
	return &Config{
		Authorization: NewAuthorizationConfig(),
		Database: database.NewPostgresConfig(),
	}
}