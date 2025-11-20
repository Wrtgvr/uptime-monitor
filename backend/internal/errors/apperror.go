package errs

import "fmt"

type AppError struct {
	Type    string   `json:"type"`
	Message string   `json:"message"`
	Err     error    `json:"-"`
	Details []string `json:"-"`
	Code    int      `json:"-"`
}

func (err *AppError) Error() string {
	return fmt.Sprintf("type=%s, code=%d, msg=%s, details=%+v err=%v", err.Type, err.Code, err.Message, err.Details, err.Err)
}
