package apperr

type (
	InvalidCredentionals string
)

func (e InvalidCredentionals) Error() string {
	return "Login or password mistmatch"
}
