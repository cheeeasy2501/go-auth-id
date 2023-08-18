package controller

import (
	"errors"
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
	CheckToken(ctx *gin.Context)
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

	srv.Response(ctx, http.StatusOK, response.NewTokenResponse(tokens.AccessToken, tokens.RefreshToken))
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

func (c *AuthorizationController) CheckToken(ctx *gin.Context) {
	uID, exist := ctx.Get("userId")
	if !exist {
		srv.ErrorResponse(ctx, http.StatusBadRequest, errors.New("user id isn't found"))
		return
	}

	userId, casted := uID.(uint64)
	if !casted {
		srv.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid user id"))
		return
	}

	request := request.NewRefreshTokenRequest(userId)

	tokens, err := c.Authorization.RefreshToken(request)
	if err != nil {
		srv.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	srv.Response(ctx, http.StatusOK, response.NewTokenResponse(tokens.AccessToken, tokens.RefreshToken))
}