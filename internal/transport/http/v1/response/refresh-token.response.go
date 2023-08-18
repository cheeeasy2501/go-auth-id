package response

type ITokenResponse interface {
	Access() string
	Refresh() string
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewTokenResponse(accessToken, refreshToken string) TokenResponse {
	return TokenResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
}

func (r *TokenResponse) Access() string {
	return r.AccessToken
}

func (r *TokenResponse) Refresh() string {
	return r.RefreshToken
}
