package errs

import "net/http"

func NewInternalError(err error, details []string) *AppError {
	return &AppError{
		Type:    TypeInternalError,
		Message: "internal server error",
		Err:     err,
		Details: details,
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFound(err error, details []string, msg string) *AppError {
	return &AppError{
		Type:    TypeNotFound,
		Message: msg,
		Err:     err,
		Details: details,
		Code:    http.StatusNotFound,
	}
}

func NewBadRequest(err error, details []string, msg string) *AppError {
	return &AppError{
		Type:    TypeBadRequest,
		Message: msg,
		Err:     err,
		Details: details,
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorized(err error, details []string, msg string) *AppError {
	return &AppError{
		Type:    TypeUnauthorized,
		Message: msg,
		Err:     err,
		Details: details,
		Code:    http.StatusUnauthorized,
	}
}

func NewForbidden(err error, details []string) *AppError {
	return &AppError{
		Type:    TypeForbidden,
		Message: "access denied",
		Err:     err,
		Details: details,
		Code:    http.StatusForbidden,
	}
}
