package app

import (
	entity  "github.com/cheeeasy2501/auth-id/entity/app"
	service "github.com/cheeeasy2501/auth-id/service/app"
)

type IAuthorizationController interface {

	// TODO: вернуть токен
	// Авторизует нового пользователя
	LoginByEmail(email, password string) error

	// Регистрирует нового пользователя
	Register(user entity.UserEntity) (entity.UserEntity, error)
}

type AuthorizationController struct {
}

func NewAuthorizationController(s *service.Services) {

}

func (authorization *AuthorizationController) LoginByEmail(email, password string) error {

	return nil
}

func (authorization *AuthorizationController) Register(user User) (User, error) {

	return user, nil
}
