package service

import (
	"context"

	"github.com/cheeeasy2501/auth-id/internal/entity"
	"gorm.io/gorm"
)

type IUserService interface {
	GetUserById(ctx context.Context, id uint64) (entity.User, error)
}

type UserService struct {
	conn *gorm.DB
}

func NewUserService(conn *gorm.DB) IUserService {
	return &UserService{
		conn: conn,
	}
}

func (s *UserService) GetUserById(ctx context.Context, id uint64) (entity.User, error) {
	u := entity.NewUser()
	res := s.conn.First(&u, id)
	if err := res.Error; err != nil {
		return u, err
	}

	return u, nil
}
