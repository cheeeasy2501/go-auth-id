package authorization

import (
	context "context"
	pb "github.com/cheeeasy2501/pb/authorization"
)

type AuthorizeGRPCServer struct {
}

func (s *AuthorizeGRPCServer) mustEmbedUnimplementedAuthServiceServer() {
}

// TODO: mock
func (s *AuthorizeGRPCServer) CheckToken(ctx context.Context, in *pb.CheckTokenRequest) (*CheckTokenResponse, error) {
	return &pb.CheckTokenResponse{
		Authorize: true,
		UserId:    1,
	}, nil
}

func (s *AuthorizeGRPCServer) GetUserInformation(ctx context.Context, in *GetUserInformationRequest) (*GetUserResponse, error) {
	return &pb.GetUserResponse{
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
