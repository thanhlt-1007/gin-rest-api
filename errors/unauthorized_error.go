package errors

type UnauthorizedError struct {
	RootError error
}

func (err UnauthorizedError) Error() string {
	return err.RootError.Error()
}
