// Package token provides JWT token generation and validation utilities.
package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
	"github.com/dawit_hopes/saas/auth/internal/domain/port/outbound"
)

var jwtSecret = []byte("")

type JWTProvider struct {
	signingKey []byte
	issuer     string
}

func NewTokenProvider(signingKey []byte, issuer string) outbound.TokenProvider {
	return &JWTProvider{
		signingKey: signingKey,
		issuer:     issuer,
	}
}

func (p *JWTProvider) GenerateToken(userID string, ttl time.Duration) (string, *appErr.AppError) {
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		Issuer:    p.issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(p.signingKey)
	if err != nil {
		return "", appErr.ErrInternalServer
	}
	return signedToken, nil
}

func (p *JWTProvider) ValidateToken(tokenStr string) (*jwt.RegisteredClaims, *appErr.AppError) {
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
		return nil, appErr.ErrInvalidToken
	}

	return claims, nil
}
