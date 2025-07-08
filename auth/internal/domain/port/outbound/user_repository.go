// Package outbound provides interfaces for outbound port operations related to users.
package outbound

import "github.com/dawit_hopes/saas/auth/internal/domain/model"

type UserRepository interface {
	Create(user *model.User) error
	GetByEmail(email string) (*model.User, error)
}
