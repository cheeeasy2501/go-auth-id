package authorization

import (
	"context"

	gen "github.com/cheeeasy2501/auth-id/gen/authorization"
	srvs "github.com/cheeeasy2501/auth-id/internal/service"
)

type AuthorizationGRPCServer struct {
	gen.UnimplementedAuthorizationServiceServer
	authorizationService srvs.IAuthorizationService
	userService          srvs.IUserService
}

func NewAuthorizationGRPCServer(
	authorizationService srvs.IAuthorizationService,
	userService srvs.IUserService,
) *AuthorizationGRPCServer {
	return &AuthorizationGRPCServer{
		authorizationService: authorizationService,
		userService:          userService,
	}
}

// TODO: mock
func (s *AuthorizationGRPCServer) CheckToken(ctx context.Context, in *gen.CheckTokenRequest) (*gen.CheckTokenResponse, error) {
	return &gen.CheckTokenResponse{
		Authorize: true,
		UserId:    1,
	}, nil
}

// TODO: почему не каститься ctx?
func (s *AuthorizationGRPCServer) GetUserById(ctx context.Context, in *gen.GetUserByIdRequest) (*gen.GetUserByIdResponse, error) {
	u, err := s.userService.GetUserById(nil, in.GetUserId())
	if err != nil {
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
