package app

type Services struct {
	authorization IAuthorizationService  
}

func NewServices(config IConfig) *Services{
	return &Services{
		authorization: NewAuthorizationService(config.GeSecretKey()),
	}
}