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

// Проверяет токен
func (m *JWTMiddleware) Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Autorization")
		if header == "" {
			server.ErrorResponse(ctx, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		parts := strings.Split(header, " ")
		if len(parts) != 2 {
			server.ErrorResponse(ctx, http.StatusUnauthorized, errors.New("Invalid token"))
			return
		}

		accessToken := parts[1]

		userId, err := m.s.ParseToken(accessToken)
		if err != nil {
			server.ErrorResponse(ctx, http.StatusUnauthorized, err)
			return
		}

		ctx.Set("userId", userId)
	}
}
