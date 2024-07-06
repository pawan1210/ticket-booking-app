package errors

type BadRequest struct {
	message string
}

func NewBadRequestError(message string) *BadRequest {
	return &BadRequest{
		message,
	}
}

func (e *BadRequest) Error() string {
	return e.message
}
