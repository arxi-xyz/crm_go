package appError

import (
	"errors"
	"net/http"
)

type AppError struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Status  int            `json:"-"`
	Meta    map[string]any `json:"meta,omitempty"`
	Err     error          `json:"-"`
}

func (e *AppError) Error() string {
	return e.Code
}

func (e *AppError) Unwrap() error { return e.Err }

func New(status int, code, message string, err error, meta map[string]any) *AppError {
	return &AppError{
		Status:  status,
		Code:    code,
		Message: message,
		Err:     err,
		Meta:    meta,
	}
}

func Validation(message string, fields map[string][]string) *AppError {
	meta := map[string]any{"fields": fields}
	return New(http.StatusBadRequest, "validation_error", message, nil, meta)
}

func NotFound(code, message string, err error) *AppError {
	return New(http.StatusNotFound, code, message, err, nil)
}

func Unauthorized(code, message string, err error) *AppError {
	return New(http.StatusUnauthorized, code, message, err, nil)
}

func Forbidden(code, message string, err error) *AppError {
	return New(http.StatusForbidden, code, message, err, nil)
}

func Conflict(code, message string, err error) *AppError {
	return New(http.StatusConflict, code, message, err, nil)
}

func Internal(err error) *AppError {
	return New(http.StatusInternalServerError, "internal_error", "خطای داخلی سرور", err, nil)
}

func AsAppError(err error) (*AppError, bool) {
	var ae *AppError
	if errors.As(err, &ae) {
		return ae, true
	}
	return nil, false
}
