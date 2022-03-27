package common

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "NOT_FOUND_ERROR"
}

type InternalError struct{}

func (e *InternalError) Error() string {
	return "INTERNAL_ERROR"
}
