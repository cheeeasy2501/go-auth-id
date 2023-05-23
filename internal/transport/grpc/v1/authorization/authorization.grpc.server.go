package authorization

import (
	"context"
	"fmt"
	_ "fmt"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gen "github.com/cheeeasy2501/auth-id/gen"
	"github.com/cheeeasy2501/auth-id/internal/apperr"
	_ "github.com/cheeeasy2501/auth-id/internal/apperr"
	srvs "github.com/cheeeasy2501/auth-id/internal/service"
	auth_pkg "github.com/cheeeasy2501/auth-id/pkg/grpc/v1/authorization"
	"google.golang.org/grpc"
)

type AuthorizationGRPCServer struct {
	gen.UnimplementedAuthorizationServiceServer
	srv                  *grpc.Server
	cfg                  auth_pkg.IConfig
	authorizationService srvs.IAuthorizationService
	userService          srvs.IUserService
}

func NewAuthorizationGRPCServer(cfg auth_pkg.IConfig, srvs *srvs.Services) (*AuthorizationGRPCServer, error) {
	return &AuthorizationGRPCServer{
		cfg:                  cfg,
		authorizationService: srvs.Authorization,
		userService:          srvs.User,
	}, nil
}

func (s *AuthorizationGRPCServer) Run(srvs *srvs.Services) error {
	s.srv = grpc.NewServer()
	// srv := NewAuthorizationGRPCServer(s.cfg, srvs)
	gen.RegisterAuthorizationServiceServer(s.srv, s) // TODO: check

	l, err := net.Listen("tcp", s.cfg.GetAddr())
	if err != nil {
		return err
	}

	go func() {
		if err := s.srv.Serve(l); err != nil {
			fmt.Printf("GRPC Server Shutdown. Error: %v\n", err)
		}
	}()

	return nil
}

func (s *AuthorizationGRPCServer) CheckToken(ctx context.Context, in *gen.CheckTokenRequest) (*gen.CheckTokenResponse, error) {
	id, err := s.authorizationService.ParseToken(in.GetToken())
	if err != nil {
		err = status.Error(codes.PermissionDenied, err.Error())
		return &gen.CheckTokenResponse{
			Authorize: false,
			UserId:    0,
		}, err
	}

	return &gen.CheckTokenResponse{
		Authorize: true,
		UserId:    id,
	}, nil
}

func (s *AuthorizationGRPCServer) GetUserById(ctx context.Context, in *gen.GetUserByIdRequest) (*gen.GetUserByIdResponse, error) {
	u, err := s.userService.GetUserById(ctx, in.GetUserId())
	if err != nil {
		err = status.Error(codes.NotFound, apperr.NewShortCustomErr("GetUserById", err).Error())
		return &gen.GetUserByIdResponse{}, err
	}

	return &gen.GetUserByIdResponse{
		ID:         u.ID,
		Avatar:     u.Avatar,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		MiddleName: u.MiddleName,
		Email:      u.Email,
	}, nil
}
