// Package token provides JWT token generation and validation utilities.
package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
)

var jwtSecret = []byte("")

type JWTProvider struct {
	signingKey []byte
	issuer     string
	ttl        time.Duration
}

func NewTokenProvide(signingKey []byte, issuer string, ttl time.Duration) *JWTProvider {
	return &JWTProvider{
		signingKey: signingKey,
		issuer:     issuer,
		ttl:        ttl,
	}
}

func (p *JWTProvider) GenerateToken(userID string, roles []string) (string, error) {

	claims := jwt.RegisteredClaims{
		Subject:   userID,
		Issuer:    p.issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(p.ttl)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(p.signingKey)
}

func (p *JWTProvider) ValidateToken(tokenStr string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, appErr.ErrInvalidSigningMethod
		}
		return p.signingKey, nil
	})
	if err != nil {
		return nil, appErr.ErrInternalServer
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, appErr.ErrInvlidToken
	}

	return claims, nil
}
