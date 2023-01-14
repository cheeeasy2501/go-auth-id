package controller

import (
	"github.com/cheeeasy2501/auth-id/internal/service"
	_ "github.com/gin-gonic/gin"
)

type Controller struct {
	Authorization IAuthorizationController
}

func NewController(s *service.Services) *Controller {
	return &Controller{
		Authorization: NewAuthorizationController(s),
	}
}