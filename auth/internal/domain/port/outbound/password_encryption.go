package outbound

import (
	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
    
    "context")



// PasswordSecurity defines the interface for password hashing and verification
type PasswordSecurity interface {
    HashPassword(ctx context.Context, password string) (string, *appErr.AppError)    
    ComparePassword(ctx context.Context, hashedPassword, password string) *appErr.AppError
}