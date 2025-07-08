// Package errors provides custom error types for application error handling.
package errors

type AppError struct {
	Code     int
	Message  string
	Internal string
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
