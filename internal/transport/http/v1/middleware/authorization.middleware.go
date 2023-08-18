package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/cheeeasy2501/auth-id/internal/service"
	"github.com/cheeeasy2501/auth-id/pkg/server"
	"github.com/gin-gonic/gin"
)

type JWTMiddleware struct {
	s service.ITokenService
}

func NewJWTMiddleware(s service.ITokenService) *JWTMiddleware {
	return &JWTMiddleware{
		s: s,
	}
}

func (m *JWTMiddleware) SplitToken(ctx *gin.Context, headerName string) (string, bool) {
	hn := "Authorization"

	if headerName != "" {
		hn = headerName
	}

	header := ctx.GetHeader(hn)
	if header == "" {
		server.ErrorResponse(ctx, http.StatusUnauthorized, errors.New("Unauthorized"))
		return "", false
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		server.ErrorResponse(ctx, http.StatusUnauthorized, errors.New("Invalid token"))
		return "", false
	}

	return parts[1], true
}

// Проверяет токен
func (m *JWTMiddleware) Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, ok := m.SplitToken(ctx, "Authorization")
		if !ok {
			return
		}

		userId, err := m.s.ParseToken(token)
		if err != nil {
			server.ErrorResponse(ctx, http.StatusUnauthorized, err)
			return
		}

		ctx.Set("userId", userId)
	}
}

// TODO: возможно стоит сделать одну токен-функцию
// func (m *JWTMiddleware) CheckRefreshToken() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		token, ok := m.SplitToken(ctx)
// 		if !ok {
// 			return
// 		}

// 		userId, err := m.s.ParseToken(token)
// 		if err != nil {
// 			server.ErrorResponse(ctx, http.StatusUnauthorized, err)
// 			return
// 		}

// 		ctx.Set("userId", userId)
// 	}
// }
