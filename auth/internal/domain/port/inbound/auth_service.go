// Package inbound provides interfaces for inbound authentication services.
package inbound

import (
	"context"

	"github.com/dawit_hopes/saas/auth/internal/domain/model"
	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"

)

type AuthService interface {
	Signup(ctx context.Context, user model.User) (model.Auth, *appErr.AppError)
	Login(ctx context.Context, email, password string) (model.Auth, *appErr.AppError)
	RefreshToken(ctx context.Context, refreshToken string) (model.Auth, *appErr.AppError)
	Me(ctx context.Context, accessToken string) (model.User, *appErr.AppError)
	// Logout(ctx context.Context, accessToken string) *appErr.AppError
}
