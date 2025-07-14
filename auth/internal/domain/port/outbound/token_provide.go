// Package outbound provides interfaces for outbound port operations related to token.
package outbound

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"

)

type TokenProvider interface {
	GenerateToken(userID string, ttl time.Duration) (string, *appErr.AppError)
	ValidateToken(token string) (claims *jwt.RegisteredClaims, err *appErr.AppError)
}
