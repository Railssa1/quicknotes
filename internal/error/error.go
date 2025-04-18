package error

type StatusError struct {
	error
	status int
}

type RepositoryError struct {
	error
}

func NewRepositoryError(err error) error {
	return RepositoryError{
		error: err,
	}
}

func (se StatusError) StatusCode() int {
	return se.status
}

func WithStatus(err error, status int) error {
	return StatusError{
		error:  err,
		status: status,
	}
}
