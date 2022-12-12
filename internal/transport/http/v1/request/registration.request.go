package request

type RegistrationRequest struct {
	Avatar     *string `json:"avatar" validate:"url"`
	Email      string `json:"email" validate:"required,email"`
	FirstName  string `json:"firstName" validate:""`
	LastName   string `json:"lastName" validate:""`
	MiddleName *string `json:"middleName" validate:""`
	Password   string `json:"password" validate:"required"`
}
