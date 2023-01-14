package middleware

import "github.com/cheeeasy2501/auth-id/internal/service"

type Middleware struct {
	Jwtm JWTMiddleware
}

func NewMiddleware(
	s service.ITokenService,
) *Middleware {
	return &Middleware{
		Jwtm: *NewJWTMiddleware(s),
	}
}
