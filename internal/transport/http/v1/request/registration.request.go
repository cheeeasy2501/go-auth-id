package request

type RegistrationRequest struct {
	Avatar     *string `json:"avatar" binding:"url"`
	Email      string  `json:"email" binding:"required,email"`
	FirstName  string  `json:"firstName" binding:"gte=2"`
	LastName   string  `json:"lastName" binding:"gte=2"`
	MiddleName *string `json:"middleName"`
	Password   string  `json:"password" binding:"required,gte=8"`
}
