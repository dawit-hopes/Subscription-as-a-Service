package inbound

import (
	"context"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
	"github.com/dawit_hopes/saas/auth/internal/domain/model"
	"github.com/dawit_hopes/saas/auth/internal/domain/port/outbound"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) (*model.User, *appErr.AppError)
	GetByID(ctx context.Context, id string) (*model.User, *appErr.AppError)
	Update(ctx context.Context, user *model.User) (*model.User, *appErr.AppError)
	Delete(ctx context.Context, id string) *appErr.AppError
	GetByEmail(ctx context.Context, email string) (*model.User, *appErr.AppError)
}
type userService struct {
	userRepo outbound.UserRepository
}

func NewUserService(userRepo outbound.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Create(ctx context.Context, user *model.User) (*model.User, *appErr.AppError) {
	if err := user.Validation(); err != nil {
		return nil, err
	}

	existUser, _ := u.userRepo.GetByEmail(ctx, user.Email)
	if existUser != nil {
		return nil, appErr.ErrEmailExists
	}
	return u.userRepo.Create(ctx, user)
}

func (u *userService) GetByID(ctx context.Context, id string) (*model.User, *appErr.AppError) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *userService) Update(ctx context.Context, user *model.User) (*model.User, *appErr.AppError) {
	existUser, err := u.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if existUser == nil {
		return nil, appErr.ErrUserNotFound
	}
	return u.userRepo.Update(ctx, user)
}

func (u *userService) Delete(ctx context.Context, id string) *appErr.AppError {
	existUser, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existUser == nil {
		return appErr.ErrUserNotFound
	}
	return u.userRepo.Delete(ctx, id)
}

func (u *userService) GetByEmail(ctx context.Context, email string) (*model.User, *appErr.AppError) {
	return u.userRepo.GetByEmail(ctx, email)
}
