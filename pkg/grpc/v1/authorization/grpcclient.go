package authorization

import (
	"context"
	"fmt"

	gen "github.com/cheeeasy2501/auth-id/gen"
	"google.golang.org/grpc"
)

type AuthorizationClient struct {
	ctx    context.Context
	conn   *grpc.ClientConn
	client gen.AuthorizationServiceClient
}

func (c *AuthorizationClient) Connect() error {
	conn, err := grpc.DialContext(c.ctx, ":1000")
	if err != nil {
		return fmt.Errorf("AuthorizationClient:Connection() - Not connected! %w", err)
	}

	c.conn = conn
	c.client = gen.NewAuthorizationServiceClient(c.conn)

	return nil
}

func (c *AuthorizationClient) CloseConnection() error {
	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("AuthorizationClient:CloseConnection() - Not closed! %w", err)
	}

	return nil
}

func (c *AuthorizationClient) CheckToken(ctx context.Context, in *gen.CheckTokenRequest) (*gen.CheckTokenResponse, error) {
	out, err := c.client.CheckToken(ctx, in)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Token has been verified. UserId = %d", out.GetUserId())

	return &gen.CheckTokenResponse{
		UserId:    out.GetUserId(),
		Authorize: out.GetAuthorize(),
	}, nil
}

func (c *AuthorizationClient) GetUserById(ctx context.Context, in *gen.GetUserByIdRequest) (*gen.GetUserByIdResponse, error) {
	out, err := c.client.GetUserById(ctx, in)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Information about user has been received. UserId = %d", out.GetID())

	return &gen.GetUserByIdResponse{
		ID:         out.GetID(),
		Avatar:     out.GetAvatar(),
		FirstName:  out.GetFirstName(),
		LastName:   out.GetLastName(),
		MiddleName: out.GetMiddleName(),
		Email:      out.GetEmail(),
		IsBanned:   out.GetIsBanned(),
	}, nil
}
