package service

import (
	cfg "github.com/cheeeasy2501/auth-id/config"
	"gorm.io/gorm"
)

type Services struct {
	Authorization IAuthorizationService
	User IUserService
	Mail IMailService
}

func NewService(config *cfg.Config, conn *gorm.DB) *Services {
	return &Services{
		Authorization: NewAuthorizationService(config.Authorization.GetSecretKey(), conn),
		User: NewUserService(conn),
		Mail: NewMailService(conn),
	}
}
