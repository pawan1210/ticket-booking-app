package user

type UserAlreadyExists struct {
	error string
}

func (e *UserAlreadyExists) Error() string {
	return e.error
}
