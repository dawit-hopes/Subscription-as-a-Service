package errors

import "net/http"

func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrEmailExists          = New(http.StatusBadRequest, "Email already exists")
	ErrInvalidCredentials   = New(http.StatusUnauthorized, "Invalid email or password")
	ErrUserNotFound         = New(http.StatusNotFound, "User not found")
	ErrInvalidEmailFormat   = New(http.StatusBadRequest, "invalid email format")
	ErrInvalidRole          = New(http.StatusBadRequest, "invalid role")
	ErrInvlidToken          = New(http.StatusBadRequest, "invalid token")
	ErrInvalidSigningMethod = New(http.StatusBadRequest, "unexpected signing method")
	ErrInternalServer       = New(http.StatusBadRequest, "Internal server error")
)
