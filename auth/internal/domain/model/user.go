// Package model contains domain models for the authentication service.
package model

import (
	"regexp"
	"strings"
	"time"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
)

type Role string

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type User struct {
	ID        string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser creates a new User instance after validating the email and role.
// Password should be hashed before calling this or inside service layer.
func NewUser(id, email, hashedPassword string, role Role) (*User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	if !isValidEmail(email) {
		return nil, appErr.ErrInvalidEmailFormat
	}

	return &User{
		ID:        id,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func (u *User) SetPassword(hashedPassword string) {
	u.Password = hashedPassword
	u.UpdatedAt = time.Now()
}
