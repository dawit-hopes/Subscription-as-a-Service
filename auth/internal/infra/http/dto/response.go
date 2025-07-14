package dto

import "github.com/dawit_hopes/saas/auth/internal/domain/model"

type AuthResponse struct {
	StatusCode int
	Token      model.Auth
}

type UserResponse struct {
	StatusCode int
	User model.User
}
