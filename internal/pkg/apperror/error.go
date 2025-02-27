package apperror

type ApplicationError struct {
	Status    int
	Message   string
	ErrorCode string
}

func (e *ApplicationError) Error() string {
	return e.Message
}

func New(status int, errCode string, err error) error {
	return &ApplicationError{
		Status:    status,
		ErrorCode: errCode,
		Message:   err.Error(),
	}
}
