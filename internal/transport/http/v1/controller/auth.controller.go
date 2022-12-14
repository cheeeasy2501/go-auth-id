package controller

import (
	"net/http"

	"github.com/cheeeasy2501/auth-id/internal/service"
	"github.com/cheeeasy2501/auth-id/internal/transport/http/v1/request"
	srv "github.com/cheeeasy2501/auth-id/pkg/server"

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
	request := new(request.LoginByEmailRequest)
	err := ctx.ShouldBindJSON(request)

	if err != nil {
		srv.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	token, err := c.Authorization.LoginByEmail(request)
	if err != nil {
		srv.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	srv.Response(ctx, http.StatusOK, gin.H{
		"token": token,
	})
}

func (c *AuthorizationController) Registration(ctx *gin.Context) {
	request := new(request.RegistrationRequest)
	err := ctx.ShouldBindJSON(request)

	if err != nil {
		srv.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err = c.Authorization.Registration(request)
	if err != nil {
		srv.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	srv.Response(ctx, http.StatusCreated, nil)
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
