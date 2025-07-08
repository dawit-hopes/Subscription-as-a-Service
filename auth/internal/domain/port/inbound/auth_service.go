// Package inbound provides interfaces for inbound authentication services.
package inbound

import "github.com/dawit_hopes/saas/auth/internal/domain/model"

type AuthService interface {
	Signup(email, password string, role model.Role) (string, error)
	Login(email, password string) (string, error)
}
