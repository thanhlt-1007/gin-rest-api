package errors

type UniqueContraintError struct {
    RootError error
    TableName string
    Fields []string
}

func (err UniqueContraintError) Error() string {
    return err.RootError.Error()
}
