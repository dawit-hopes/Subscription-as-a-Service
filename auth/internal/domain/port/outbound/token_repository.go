package outbound

import (
	"context"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
	"github.com/dawit_hopes/saas/auth/internal/domain/model"
)

type RefreshTokenRepository interface {
	Create(ctx context.Context, token model.RefreshToken) *appErr.AppError
	GetByUserID(ctx context.Context, userID string) (*model.RefreshToken, *appErr.AppError)
	Delete(ctx context.Context, userID string) *appErr.AppError
	Update(ctx context.Context, token model.RefreshToken) *appErr.AppError
	RevokeByUserID(ctx context.Context, userID string) *appErr.AppError
}
