package model

import (
	"time"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
)

type Auth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	Revoked   bool      `json:"revoked"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (r *RefreshToken) Validation() *appErr.AppError {
	if r.UserID == "" {
		return appErr.ErrInvalidUserID
	}
	if r.Token == "" {
		return appErr.ErrInvalidToken
	}
	return nil
}
