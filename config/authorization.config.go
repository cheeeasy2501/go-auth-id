package app

import "os"

type IAuthorizationConfig interface {
	GetSecretKey() string
}

type AuthorizationConfig struct {
	secretKey string
}

func NewAuthorizationConfig() *AuthorizationConfig {
	return &AuthorizationConfig{
		secretKey: os.Getenv("SECRET_KEY"),
	}
}

func (config *AuthorizationConfig) GetSecretKey() string {
	return config.secretKey
}