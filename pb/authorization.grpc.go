package pb

import (
	context "context"

	srvs "github.com/cheeeasy2501/internal/service"
)

type AuthorizationGRPCServer struct {
	authorizationService srvs.IAuthorizationService
}

func NewAUthorizationGRPCServer(authorizationService srvs.IAuthorizationService) *AuthorizationGRPCServer {
	return &AuthorizationGRPCServer{
		authorizationService: authorizationService,
	}
}

// mustEmbedUnimplementedAuthorizationServiceServer implements AuthorizationServiceServer
func (s *AuthorizationGRPCServer) mustEmbedUnimplementedAuthorizationServiceServer() {
}

// TODO: mock
func (s *AuthorizationGRPCServer) CheckToken(ctx context.Context, in *CheckTokenRequest) (*CheckTokenResponse, error) {
	return &CheckTokenResponse{
		Authorize: true,
		UserId:    1,
	}, nil
}

func (s *AuthorizationGRPCServer) GetUserInformation(ctx context.Context, in *GetUserInformationRequest) (*GetUserResponse, error) {
	return &GetUserResponse{
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