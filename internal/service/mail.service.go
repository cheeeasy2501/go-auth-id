package service

import "gorm.io/gorm"

type IMailService interface {
	Send() (bool, error)
}

type MailService struct {
	conn *gorm.DB
}

func NewMailService(conn *gorm.DB) IMailService {
	return &MailService{
		conn: conn,
	}
}

func (s *MailService) Send() (bool, error) {
	return true, nil
}
