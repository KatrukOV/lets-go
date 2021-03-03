package usr

import "errors"

var (
	ErrInvalidCredentials = errors.New("models: invalid user credentials")
	ErrUserAlreadyExist   = &ErrorUserAlreadyExist{}
)

type ErrorUserAlreadyExist struct {
}

func (r ErrorUserAlreadyExist) Error() string {
	return "duplicate user"
}
