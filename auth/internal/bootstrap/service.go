package bootstrap

import (
	t "github.com/dawit_hopes/saas/auth/internal/domain/port/inbound/token"
	user "github.com/dawit_hopes/saas/auth/internal/domain/port/inbound/user"

	"github.com/dawit_hopes/saas/auth/internal/domain/port/outbound"
	"github.com/dawit_hopes/saas/auth/internal/infra/security"
	"github.com/dawit_hopes/saas/auth/internal/infra/token"
)

type Service struct {
	TokenProvider    outbound.TokenProvider
	PasswordSecurity outbound.PasswordSecurity
	UserService      user.UserService
	TokenService     t.RefreshTokenService
}

func InitServices(signingKey []byte, issuer string, cost int, persistance Persistance) Service {
	return Service{
		TokenProvider:    token.NewTokenProvider(signingKey, issuer),
		PasswordSecurity: security.NewBcryptPasswordSecurity(cost),
		UserService:      user.NewUserService(persistance.UserRepository),
		TokenService:     t.NewRefreshTokenService(persistance.TokenRepositoy),
	}
}
