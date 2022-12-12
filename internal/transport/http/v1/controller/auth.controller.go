package controller

import (
	"net/http"

	"github.com/cheeeasy2501/auth-id/internal/service"
	"github.com/cheeeasy2501/auth-id/internal/transport/http/v1/request"
	"github.com/gin-gonic/gin"
)

type IAuthorizationController interface {
	LoginByEmail(ctx *gin.Context)
	Registration(ctx *gin.Context)
	RegisterRoutes(group *gin.RouterGroup)
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

func (c *AuthorizationController) Registration(ctx *gin.Context) {
	request := new(request.RegistrationRequest)
	err := ctx.ShouldBindJSON(request)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		ctx.Abort()
	}

	err = c.Authorization.Registration(request)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		ctx.Abort()
	}

	ctx.JSON(http.StatusCreated, nil)
	return
}

func (c *AuthorizationController) RegisterRoutes(group *gin.RouterGroup) {
	group.POST("/login", c.LoginByEmail)
	group.POST("/registration", c.Registration)
}

// func (c *AuthorizationController) LoginByEmail(email, password string) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		c.Authorization.LoginByEmail(email, password)
// 	}
// }

// func (c *AuthorizationController) Register(user entity.User) (entity.User, error) {

// 	return user, nil
// }
