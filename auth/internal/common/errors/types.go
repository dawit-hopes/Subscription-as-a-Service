// Package errors provides custom error types for application error handling.
package errors

type AppError struct {
	Code       int    // HTTP status code
	Message    string // user-facing message
	Internal   string // optional internal/debug message
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) FullError() string {
	if e.Internal != "" {
		return e.Message + ": " + e.Internal
	}
	return e.Message
}
