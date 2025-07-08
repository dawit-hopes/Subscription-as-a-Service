package errors

import "net/http"

func New(code int, message string, internal string) *AppError {
	return &AppError{
		Code:     code,
		Message:  message,
		Internal: internal,
	}
}

func BadRequest(msg string) *AppError {
	return New(http.StatusBadRequest, msg, msg)
}

func Unauthorized(msg string) *AppError {
	return New(http.StatusUnauthorized, msg, msg)
}

func Forbidden(msg string) *AppError {
	return New(http.StatusForbidden, msg, msg)
}

func NotFound(msg string) *AppError {
	return New(http.StatusNotFound, msg, msg)
}

func Internal(err error) *AppError {
	return New(http.StatusInternalServerError, "Internal Server Error", err.Error())
}

var (
	ErrEmailExists = New(
		http.StatusBadRequest,
		"Email already exists",
		"attempt to register duplicate",
	)

	ErrInvalidCredentials = New(
		http.StatusUnauthorized,
		"Invalid email or password",
		"login failed due to mismatched credentials",
	)

	ErrUserNotFound = New(
		http.StatusNotFound,
		"User not found",
		"user not found in database",
	)
)
