package request

type RefreshTokenRequest struct {
	// RefreshToken string `json:"refreshToken" header:"Authorization" binding:"required"`
	UserId uint64 `json:"refreshToken" header:"Authorization" binding:"required"`
}

func NewRefreshTokenRequest(userId uint64)* RefreshTokenRequest {
	return &RefreshTokenRequest{
		UserId: userId,
	}
}
