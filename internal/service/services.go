package service

import (
	cfg "github.com/cheeeasy2501/auth-id/config"
	"gorm.io/gorm"
)

type Services struct {
	Authorization IAuthorizationService
}

func NewService(config *cfg.Config, conn *gorm.DB) *Services {
	return &Services{
		Authorization: NewAuthorizationService(config.Authorization.GetSecretKey(), conn),
	}
}
