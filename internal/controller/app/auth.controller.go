package app

import (
	"net/http"

	service "github.com/cheeeasy2501/auth-id/internal/service/app"
	"github.com/gin-gonic/gin"
)

type IAuthorizationController interface {

	// TODO: вернуть токен
	// Авторизует нового пользователя
	// LoginByEmail(email, password string) error

	// Регистрирует нового пользователя
	// Register(user entity.UserEntity) (entity.UserEntity, error)
	LoginByEmail(ctx *gin.Context)
	// Register(ctx *gin.Context)
}

type AuthorizationController struct {
	Authorization service.IAuthorizationService
}

func NewAuthorizationController(s *service.Services) *AuthorizationController {
	return &AuthorizationController{
		Authorization: s.Authorization,
	}
}

func (c *AuthorizationController) LoginByEmail(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"token": "123",
	})
}

// func (c *AuthorizationController) LoginByEmail(email, password string) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		c.Authorization.LoginByEmail(email, password)
// 	}
// }

// func (c *AuthorizationController) Register(user entity.UserEntity) (entity.UserEntity, error) {

// 	return user, nil
// }
