package entity

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func NewTokens() *Tokens {
	return &Tokens{}
}
