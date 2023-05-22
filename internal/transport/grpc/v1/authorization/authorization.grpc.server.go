package authorization

import (
	"context"
	_ "fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gen "github.com/cheeeasy2501/auth-id/gen"
	"github.com/cheeeasy2501/auth-id/internal/apperr"
	_ "github.com/cheeeasy2501/auth-id/internal/apperr"
	srvs "github.com/cheeeasy2501/auth-id/internal/service"
)

type AuthorizationGRPCServer struct {
	gen.UnimplementedAuthorizationServiceServer
	authorizationService srvs.IAuthorizationService
	userService          srvs.IUserService
}

func NewAuthorizationGRPCServer(s *srvs.Services) *AuthorizationGRPCServer {
	return &AuthorizationGRPCServer{
		authorizationService: s.Authorization,
		userService:          s.User,
	}
}

// TODO: mock
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

// type AuthorizationGRPCClient struct {

// }

// func (c *AuthorizationGRPCClient) CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error) {

// }

// func (c *AuthorizationGRPCClient) GetUserInformation(ctx context.Context, in *GetUserInformationRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {

// }
