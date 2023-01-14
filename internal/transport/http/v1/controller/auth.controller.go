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
	RefreshTokens(ctx *gin.Context)
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

	tokens, err := c.Authorization.LoginByEmail(request)
	if err != nil {
		srv.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	srv.Response(ctx, http.StatusOK, gin.H{
		"accessToken":  tokens.AccessToken,
		"refreshToken": tokens.RefreshToken,
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

func (c *AuthorizationController) RefreshTokens(ctx *gin.Context) {
	request := new(request.RefreshTokens)
	err := ctx.ShouldBindJSON(request)
	if err != nil {
		srv.ErrorResponse(ctx, http.StatusBadRequest, nil)
		return
	}

	tokens, err := c.Authorization.RefreshTokens(request)
	if err != nil {
		srv.ErrorResponse(ctx, http.StatusBadRequest, err)
	}

	srv.Response(ctx, http.StatusCreated, gin.H{
		"accessToken":  tokens.AccessToken,
		"refreshToken": tokens.RefreshToken,
	})
}

// func (c *AuthorizationController) RegisterRoutes(group *gin.RouterGroup) {
// 	group.POST("/login", c.LoginByEmail)
// 	group.POST("/registration", c.Registration)
// 	group.Use()
// }
