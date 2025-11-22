package errs

import "net/http"

// 4XX
func NewNotFound(err error) *AppError {
	return &AppError{
		Type:    TypeNotFound,
		Message: "Not Found",
		Err:     err,
		Details: make([]string, 0),
		Code:    http.StatusNotFound,
	}
}

func NewBadRequest(err error) *AppError {
	return &AppError{
		Type:    TypeBadRequest,
		Message: "Bad Request",
		Err:     err,
		Details: make([]string, 0),
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorized(err error) *AppError {
	return &AppError{
		Type:    TypeUnauthorized,
		Message: "Unauthorized",
		Err:     err,
		Details: make([]string, 0),
		Code:    http.StatusUnauthorized,
	}
}

func NewForbidden(err error) *AppError {
	return &AppError{
		Type:    TypeForbidden,
		Message: "Forbidden",
		Err:     err,
		Details: make([]string, 0),
		Code:    http.StatusForbidden,
	}
}

func NewConflict(err error) *AppError {
	return &AppError{
		Type:    TypeConflict,
		Message: "Conflict",
		Err:     err,
		Details: make([]string, 0),
		Code:    http.StatusConflict,
	}
}

func NewTooManyRequests(err error) *AppError {
	return &AppError{
		Type:    TypeTooManyRequests,
		Message: "Too Many Requests",
		Err:     err,
		Details: make([]string, 0),
		Code:    http.StatusTooManyRequests,
	}
}

// 5XX
func NewInternalError(err error) *AppError {
	return &AppError{
		Type:    TypeInternalError,
		Message: "Internal Server Error",
		Err:     err,
		Details: make([]string, 0),
		Code:    http.StatusInternalServerError,
	}
}
