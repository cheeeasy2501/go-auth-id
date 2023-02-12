package controller

import (
	"net/http"

	"github.com/cheeeasy2501/auth-id/internal/service"
	"github.com/cheeeasy2501/auth-id/internal/transport/http/v1/request"
	"github.com/cheeeasy2501/auth-id/internal/transport/http/v1/response"
	srv "github.com/cheeeasy2501/auth-id/pkg/server"

	"github.com/gin-gonic/gin"
)

type IAuthorizationController interface {
	LoginByEmail(ctx *gin.Context)
	Registration(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
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

func (c *AuthorizationController) RefreshToken(ctx *gin.Context) {
	userIdString, exist := ctx.Get("userId")
	if exist == false {
		srv.ErrorResponse(ctx, http.StatusBadRequest, nil)
		return
	}

	userId, casted := userIdString.(uint)
	if casted == false {
		srv.ErrorResponse(ctx, http.StatusBadRequest, nil)
		return
	}

	request := request.NewRefreshTokenRequest(uint64(userId))

	tokens, err := c.Authorization.RefreshToken(request)
	if err != nil {
		srv.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	srv.Response(ctx, http.StatusCreated, response.NewRefreshTokenResponse(tokens.AccessToken, tokens.RefreshToken))
}

// func (c *AuthorizationController) RegisterRoutes(group *gin.RouterGroup) {
// 	group.POST("/login", c.LoginByEmail)
// 	group.POST("/registration", c.Registration)
// 	group.Use()
// }
