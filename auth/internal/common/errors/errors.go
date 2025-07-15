package errors

import "net/http"

func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrEmailExists               = New(http.StatusBadRequest, "Email already exists")
	ErrInvalidCredentials        = New(http.StatusUnauthorized, "Invalid credentials")
	ErrUserNotFound              = New(http.StatusNotFound, "User not found")
	ErrInvalidEmailFormat        = New(http.StatusBadRequest, "invalid email format")
	ErrInvalidSigningMethod      = New(http.StatusBadRequest, "unexpected signing method")
	ErrInternalServer            = New(http.StatusBadRequest, "Internal server error")
	ErrGeneralDatabaseInsert     = New(http.StatusInternalServerError, "General database insert error")
	ErrGeneralDatabaseUpdate     = New(http.StatusInternalServerError, "General database update error")
	ErrGeneralDatabaseDelete     = New(http.StatusInternalServerError, "General database delete error")
	ErrGeneralDatabaseQuery      = New(http.StatusInternalServerError, "General database query error")
	ErrGeneralDatabaseConnection = New(http.StatusInternalServerError, "General database connection error")
	ErrInvalidToken              = New(http.StatusBadRequest, "invalid token")
	ErrInvalidJSONPayload        = New(http.StatusBadRequest, "invalid JSON payload")
	ErrInvalidID                 = New(http.StatusBadRequest, "Invalid id format")
	ErrInvalidUserID             = New(http.StatusBadRequest, "Invalid user ID")
	ErrDocumentNotFound          = New(http.StatusNotFound, "Document not found")
	ErrTokenAlreadyExists        = New(http.StatusBadRequest, "token already exists for user")
	ErrFailedToRevokToken        = New(http.StatusInternalServerError, "failed to revoke token")
)
