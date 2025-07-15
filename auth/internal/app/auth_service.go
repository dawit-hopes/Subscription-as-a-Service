// Package app provides authentication services and related business logic.
package app

import (
	"context"
	"time"

	"github.com/dawit_hopes/saas/auth/internal/domain/model"
	auth "github.com/dawit_hopes/saas/auth/internal/domain/port/inbound/auth"
	token "github.com/dawit_hopes/saas/auth/internal/domain/port/inbound/token"
	user "github.com/dawit_hopes/saas/auth/internal/domain/port/inbound/user"

	"github.com/dawit_hopes/saas/auth/internal/domain/port/outbound"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
)

const (
	AccessTokenTTL  = 60 * time.Hour
	RefreshTokenTTL = 30 * 24 * time.Hour
)

type authService struct {
	userService     user.UserService
	tokenService    token.RefreshTokenService
	tokenProvider   outbound.TokenProvider
	passwordService outbound.PasswordSecurity
}

func NewAuthService(userService user.UserService, tokenProvider outbound.TokenProvider, passwordService outbound.PasswordSecurity, tokenService token.RefreshTokenService) auth.AuthService {
	return &authService{
		userService:     userService,
		tokenProvider:   tokenProvider,
		passwordService: passwordService,
		tokenService:    tokenService,
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

func (a *authService) newRefreshToken(userID, token string) model.RefreshToken {
	return model.RefreshToken{
		UserID:    userID,
		Token:     token,
		Revoked:   false,
		ExpiresAt: time.Now().Add(RefreshTokenTTL),
	}
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
	newUser, err := a.userService.Create(ctx, &user)
	if err != nil {
		return model.Auth{}, err
	}
	token, err := a.generateToken(newUser.ID)
	if err != nil {
		return model.Auth{}, err
	}

	if err := a.tokenService.Create(ctx, a.newRefreshToken(newUser.ID, token.RefreshToken)); err != nil {
		return model.Auth{}, err
	}
	return token, nil
}

func (a *authService) Login(ctx context.Context, email, password string) (model.Auth, *appErr.AppError) {
	existUser, err := a.userService.GetByEmail(ctx, email)
	if err != nil {
		return model.Auth{}, err
	}

	if err := a.passwordService.ComparePassword(ctx, existUser.Password, password); err != nil {
		return model.Auth{}, appErr.ErrInvalidCredentials
	}

	authTokens, err := a.generateToken(existUser.ID)
	if err != nil {
		return model.Auth{}, appErr.ErrInternalServer
	}

	if err := a.tokenService.Update(ctx, a.newRefreshToken(existUser.ID, authTokens.RefreshToken)); err != nil {
		return model.Auth{}, err
	}

	return authTokens, nil
}

func (a *authService) RefreshToken(ctx context.Context, refreshToken string) (model.Auth, *appErr.AppError) {
	userID, err := a.getUserID(refreshToken)
	if err != nil {
		return model.Auth{}, err
	}

	data, err := a.tokenService.GetByUserID(ctx, userID)
	if err != nil {
		return model.Auth{}, err
	}

	if data.Token != refreshToken || data.Revoked {
		return model.Auth{}, appErr.ErrInvalidToken
	}

	if time.Now().After(data.ExpiresAt) {
		return model.Auth{}, appErr.ErrInvalidToken
	}

	authTokens, err := a.generateToken(userID)
	if err != nil {
		return model.Auth{}, appErr.ErrInternalServer
	}

	if err := a.tokenService.Update(ctx, model.RefreshToken{
		UserID:    userID,
		Token:     authTokens.RefreshToken,
		Revoked:   false,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}); err != nil {
		return model.Auth{}, err
	}

	return authTokens, nil
}

func (a *authService) Me(ctx context.Context, userID string) (model.User, *appErr.AppError) {
	token, err := a.tokenService.GetByUserID(ctx, userID)
	if err != nil {
		return model.User{}, err
	}

	if token.Revoked {
		return model.User{}, appErr.ErrInvalidToken
	}

	user, err := a.userService.GetByID(ctx, userID)
	if err != nil {
		return model.User{}, err
	}

	return *user, nil
}

func (a *authService) Logout(ctx context.Context, accessToken string) *appErr.AppError {
	userID, err := a.getUserID(accessToken)
	if err != nil {
		return err
	}

	if err := a.tokenService.RevokeByUserID(ctx, userID); err != nil {
		return err
	}
	return nil
}
