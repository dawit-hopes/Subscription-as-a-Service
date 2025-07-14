// Package outbound provides interfaces for outbound port operations related to users.
package outbound

import (
	"context"

	"github.com/dawit_hopes/saas/auth/internal/domain/model"
	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"

)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, *appErr.AppError)
	GetByID(ctx context.Context, id string) (*model.User, *appErr.AppError)
	Update(ctx context.Context, user *model.User) (*model.User, *appErr.AppError)
	Delete(ctx context.Context, id string) *appErr.AppError
	GetByEmail(ctx context.Context, email string) (*model.User, *appErr.AppError)
}
