// Package model contains domain models for the authentication service.
package model

import (
	"regexp"
	"time"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
)

type Role string

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (u *User) isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func (u *User) SetPassword(hashedPassword string) {
	u.Password = hashedPassword
	u.UpdatedAt = time.Now()
}

func (u *User) Validation() *appErr.AppError {
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
