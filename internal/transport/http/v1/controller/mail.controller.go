package controller

import (
	"github.com/cheeeasy2501/auth-id/internal/service"
)

type IMailController interface {
	Send() (bool, error)
}

type MailController struct {
	Mail service.IMailService
}

func NewMailController(s *service.Services) IMailController {
	return &MailController{
		Mail: s.Mail,
	}
}

func (c *MailController) Send() (bool, error) {
	return true, nil
}
