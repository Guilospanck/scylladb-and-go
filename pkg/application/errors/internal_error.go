package errors

type InternalError struct {
	Message string
}

func (err InternalError) Error() string {
	return err.Message
}

func NewInternalError(message string) InternalError {
	return InternalError{
		Message: message,
	}
}
