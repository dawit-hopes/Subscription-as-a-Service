package bootstrap

import (
	"github.com/dawit_hopes/saas/auth/internal/domain/port/inbound"
	"github.com/dawit_hopes/saas/auth/internal/domain/port/outbound"
	"github.com/dawit_hopes/saas/auth/internal/infra/security"
	"github.com/dawit_hopes/saas/auth/internal/infra/token"
)

type Service struct {
	TokenProvider    outbound.TokenProvider
	PasswordSecurity outbound.PasswordSecurity
	UserService      inbound.UserUservice
}

func InitServices(signingKey []byte, issuer string, cost int, persistance Persistance) Service {
	return Service{
		TokenProvider:    token.NewTokenProvider(signingKey, issuer),
		PasswordSecurity: security.NewBcryptPasswordSecurity(cost),
		UserService:      inbound.NewUserService(persistance.UserRepository),
	}
}
