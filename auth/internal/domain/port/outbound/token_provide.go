// Package outbound provides interfaces for outbound port operations related to token.
package outbound

type TokenProvider interface {
	GenerateToken(userID string, role string) (string, error)
	ValidateToken(token string) (userID string, role string, err error)
}
