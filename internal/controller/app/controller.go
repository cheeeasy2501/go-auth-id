package app

import (
	service "github.com/cheeeasy2501/auth-id/internal/service/app"
)

type Controller struct {
	Authorization IAuthorizationController
}

func NewController(s *service.Services) *Controller {
	return &Controller{
		Authorization: NewAuthorizationController(s),
	}
}
