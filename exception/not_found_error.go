package exception

type NotFoundError struct {
	s string
}

func NewNotFoundError(text string) error {
	return NotFoundError{s: text}
}

func (error NotFoundError) Error() string {
	return error.s
}