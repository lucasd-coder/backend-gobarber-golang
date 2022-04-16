package errs

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r *AppError) Error() string {
	return r.Message
}
