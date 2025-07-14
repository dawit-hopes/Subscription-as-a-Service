package security

import (
	"context"
	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
	"golang.org/x/crypto/bcrypt"
)

// BcryptPasswordSecurity implements the PasswordSecurity interface using bcrypt
type BcryptPasswordSecurity struct {
	cost int // bcrypt cost factor
}

// NewBcryptPasswordSecurity creates a new BcryptPasswordSecurity instance
func NewBcryptPasswordSecurity(cost int) *BcryptPasswordSecurity {
	if cost == 0 {
		cost = bcrypt.DefaultCost // Use default cost if none provided
	}
	return &BcryptPasswordSecurity{cost: cost}
}

// HashPassword generates a hashed password using bcrypt
func (s *BcryptPasswordSecurity) HashPassword(ctx context.Context, password string) (string, *appErr.AppError) {
	// Check for context cancellation
	select {
	case <-ctx.Done():
		return "", appErr.ErrInternalServer
	default:
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), s.cost)
	if err != nil {
		return "", appErr.ErrInternalServer
	}
	return string(hashedPassword), nil
}

// ComparePassword verifies a plaintext password against a hashed password
func (s *BcryptPasswordSecurity) ComparePassword(ctx context.Context, hashedPassword, password string) *appErr.AppError {
	// Check for context cancellation
	select {
	case <-ctx.Done():
		return appErr.ErrInternalServer
	default:
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return appErr.ErrInternalServer
	}

	return nil
}
