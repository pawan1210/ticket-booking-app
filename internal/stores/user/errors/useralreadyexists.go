package errors

type UserAlreadyExists struct {
	error string
}

func (e *UserAlreadyExists) Error() string {
	return e.error
}

func NewUserAlreadyExistsError(message string) *UserAlreadyExists {
	return &UserAlreadyExists{
		error: message,
	}
}
