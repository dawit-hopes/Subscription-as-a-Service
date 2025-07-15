package token

import (
	"context"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
	"github.com/dawit_hopes/saas/auth/internal/domain/model"
	"github.com/dawit_hopes/saas/auth/internal/domain/port/outbound"
)

type RefreshTokenService interface {
	Create(ctx context.Context, token model.RefreshToken) *appErr.AppError
	GetByUserID(ctx context.Context, userID string) (*model.RefreshToken, *appErr.AppError)
	Delete(ctx context.Context, userID string) *appErr.AppError
	Update(ctx context.Context, token model.RefreshToken) *appErr.AppError
	RevokeByUserID(ctx context.Context, userID string) *appErr.AppError
}

type tokenService struct {
	tokenRepo outbound.RefreshTokenRepository
}

func NewRefreshTokenService(tokenRepo outbound.RefreshTokenRepository) RefreshTokenService {
	return &tokenService{
		tokenRepo: tokenRepo,
	}
}

func (t *tokenService) Create(ctx context.Context, token model.RefreshToken) *appErr.AppError {
	return t.tokenRepo.Create(ctx, token)
}

func (t *tokenService) GetByUserID(ctx context.Context, userID string) (*model.RefreshToken, *appErr.AppError) {
	return t.tokenRepo.GetByUserID(ctx, userID)
}

func (t *tokenService) Delete(ctx context.Context, userID string) *appErr.AppError {
	return t.tokenRepo.Delete(ctx, userID)
}

func (t *tokenService) Update(ctx context.Context, token model.RefreshToken) *appErr.AppError {
	return t.tokenRepo.Update(ctx, token)
}

func (t *tokenService) RevokeByUserID(ctx context.Context, userID string) *appErr.AppError {
	return t.tokenRepo.RevokeByUserID(ctx, userID)
}
