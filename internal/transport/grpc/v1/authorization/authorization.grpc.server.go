package authorization

import (
	context "context"

	gen "github.com/cheeeasy2501/auth-id/gen/authorization"

	srvs "github.com/cheeeasy2501/auth-id/internal/service"
)

type AuthorizationGRPCServer struct {
	gen.UnimplementedAuthorizationServiceServer
	authorizationService srvs.IAuthorizationService
}

func NewAUthorizationGRPCServer(authorizationService srvs.IAuthorizationService) *AuthorizationGRPCServer {
	return &AuthorizationGRPCServer{
		authorizationService: authorizationService,
	}
}

// TODO: mock
func (s *AuthorizationGRPCServer) CheckToken(ctx context.Context, in *gen.CheckTokenRequest) (*gen.CheckTokenResponse, error) {
	return &gen.CheckTokenResponse{
		Authorize: true,
		UserId:    1,
	}, nil
}

func (s *AuthorizationGRPCServer) GetUserInformation(ctx context.Context, in *gen.GetUserInformationRequest) (*gen.GetUserResponse, error) {
	return &gen.GetUserResponse{
		Avatar:     "test",
		FirstName:  "test",
		LastName:   "test",
		MiddleName: "test",
		Email:      "test",
	}, nil
}

// type AuthorizationGRPCClient struct {

// }

// func (c *AuthorizationGRPCClient) CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error) {

// }

// func (c *AuthorizationGRPCClient) GetUserInformation(ctx context.Context, in *GetUserInformationRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {

// }
