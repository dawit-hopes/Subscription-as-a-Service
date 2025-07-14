package dto

import (
	"regexp"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type SignUp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Location string `json:"location"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

func (u *SignUp) isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func (u *SignUp) Validation() *appErr.AppError {
	if u.Email == "" {
		return appErr.ErrInvalidEmailFormat
	}
	if !u.isValidEmail(u.Email) {
		return appErr.ErrInvalidEmailFormat
	}
	if len(u.Password) < 6 {
		return appErr.ErrInvalidCredentials
	}
	return nil
}
