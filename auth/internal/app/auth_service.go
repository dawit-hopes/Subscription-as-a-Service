// Package app provides authentication services and related business logic.
package app

import (
	"context"
	"time"

	"github.com/dawit_hopes/saas/auth/internal/domain/model"
	"github.com/dawit_hopes/saas/auth/internal/domain/port/inbound"
	"github.com/dawit_hopes/saas/auth/internal/domain/port/outbound"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
)

type authService struct {
	userApp         inbound.UserUservice
	tokenProvider   outbound.TokenProvider
	passwordService outbound.PasswordSecurity
}

func NewAuthService(userApp inbound.UserUservice, tokenProvider outbound.TokenProvider, passwordService outbound.PasswordSecurity) inbound.AuthService {
	return &authService{
		userApp:         userApp,
		tokenProvider:   tokenProvider,
		passwordService: passwordService,
	}
}

func (a *authService) getUserID(token string) (string, *appErr.AppError) {
	claims, err := a.tokenProvider.ValidateToken(token)
	if err != nil {
		return "", appErr.ErrInvalidToken
	}

	var userID string
	if claims != nil && claims.Subject != "" {
		userID = claims.Subject
	} else {
		return "", appErr.ErrInvalidToken
	}

	return userID, nil
}

func (a *authService) generateToken(userID string) (model.Auth, *appErr.AppError) {
	accessToken, err := a.tokenProvider.GenerateToken(userID, (60 * time.Hour))
	if err != nil {
		return model.Auth{}, err
	}

	refreshToken, err := a.tokenProvider.GenerateToken(userID, (30 * 24 * time.Hour))
	if err != nil {
		return model.Auth{}, err
	}

	return model.Auth{AccessToken: accessToken, RefreshToken: refreshToken}, nil

}

func (a *authService) Signup(ctx context.Context, user model.User) (model.Auth, *appErr.AppError) {
	if err := user.Validation(); err != nil {
		return model.Auth{}, err
	}
	hashedPassword, err := a.passwordService.HashPassword(ctx, user.Password)
	if err != nil {
		return model.Auth{}, err
	}
	user.Password = hashedPassword
	newUser, err := a.userApp.Create(ctx, &user)
	if err != nil {
		return model.Auth{}, err
	}
	token, err := a.generateToken(newUser.ID)
	if err != nil {
		return model.Auth{}, err
	}
	return token, nil
}

func (a *authService) Login(ctx context.Context, email, password string) (model.Auth, *appErr.AppError) {
	existUser, err := a.userApp.GetByEmail(ctx, email)
	if err != nil {
		return model.Auth{}, err
	}

	if err := a.passwordService.ComparePassword(ctx, existUser.Password, password); err != nil {
		return model.Auth{}, appErr.ErrInvalidCredentials
	}

	token, err := a.generateToken(existUser.ID)
	if err != nil {
		return model.Auth{}, appErr.ErrInternalServer
	}

	return token, nil
}

func (a *authService) RefreshToken(ctx context.Context, refreshToken string) (model.Auth, *appErr.AppError) {
	userID, err := a.getUserID(refreshToken)
	if err != nil {
		return model.Auth{}, err
	}

	token, err := a.generateToken(userID)
	if err != nil {
		return model.Auth{}, appErr.ErrInternalServer
	}

	return token, nil
}

func (a *authService) Me(ctx context.Context, accessToken string) (model.User, *appErr.AppError) {
	userID, err := a.getUserID(accessToken)
	if err != nil {
		return model.User{}, err
	}

	user, err := a.userApp.GetByID(ctx, userID)
	if err != nil {
		return model.User{}, appErr.ErrInternalServer
	}

	return *user, nil
}
