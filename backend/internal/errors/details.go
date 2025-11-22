package errs

func (err *AppError) WithDetail(s string) *AppError {
	err.Details = append(err.Details, s)
	return err
}
