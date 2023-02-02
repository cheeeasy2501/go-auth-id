package pb

import (
	context "context"
)

type GRPCServer struct {
}

func (s *GRPCServer) mustEmbedUnimplementedAuthServiceServer() {
}

// TODO: mock
func (s *GRPCServer) CheckToken(ctx context.Context, in *CheckTokenRequest) (*CheckTokenResponse, error) {
	return &CheckTokenResponse{
		Authorize: true,
		UserId:    1,
	}, nil
}

func (s *GRPCServer) GetUserInformation(ctx context.Context, in *GetUserInformationRequest) (*GetUserResponse, error) {
	return &GetUserResponse{
		Avatar:     "test",
		FirstName:  "test",
		LastName:   "test",
		MiddleName: "test",
		Email:      "test",
	}, nil
}
