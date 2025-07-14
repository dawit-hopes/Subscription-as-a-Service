package inbound

import (
	"context"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
	"github.com/dawit_hopes/saas/auth/internal/domain/model"
	"github.com/dawit_hopes/saas/auth/internal/domain/port/outbound"
)

type UserUservice interface {
	Create(ctx context.Context, user *model.User) (*model.User, *appErr.AppError)
	GetByID(ctx context.Context, id string) (*model.User, *appErr.AppError)
	Update(ctx context.Context, user *model.User) (*model.User, *appErr.AppError)
	Delete(ctx context.Context, id string) *appErr.AppError
	GetByEmail(ctx context.Context, email string) (*model.User, *appErr.AppError)
}
type userApplication struct {
	userApplication outbound.UserRepository
}

func NewUserService(userApp outbound.UserRepository) UserUservice {
	return &userApplication{
		userApplication: userApp,
	}
}

func (u *userApplication) Create(ctx context.Context, user *model.User) (*model.User, *appErr.AppError) {
	if err := user.Validation(); err != nil {
		return nil, err
	}

	existUser, err := u.userApplication.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if existUser != nil {
		return nil, appErr.ErrEmailExists
	}
	return u.userApplication.Create(ctx, user)
}

func (u *userApplication) GetByID(ctx context.Context, id string) (*model.User, *appErr.AppError) {
	return u.userApplication.GetByID(ctx, id)
}

func (u *userApplication) Update(ctx context.Context, user *model.User) (*model.User, *appErr.AppError) {
	existUser, err := u.userApplication.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if existUser == nil {
		return nil, appErr.ErrUserNotFound
	}
	return u.userApplication.Update(ctx, user)
}

func (u *userApplication) Delete(ctx context.Context, id string) *appErr.AppError {
	existUser, err := u.userApplication.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existUser == nil {
		return appErr.ErrUserNotFound
	}
	return u.userApplication.Delete(ctx, id)
}

func (u *userApplication) GetByEmail(ctx context.Context, email string) (*model.User, *appErr.AppError) {
	return u.userApplication.GetByEmail(ctx, email)
}
