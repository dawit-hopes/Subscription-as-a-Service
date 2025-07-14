package bootstrap

import (
	"github.com/dawit_hopes/saas/auth/internal/app"
	"github.com/dawit_hopes/saas/auth/internal/domain/port/inbound"
)

type Application struct {
	AuthApplication inbound.AuthService
}

func InitApplication(service Service) Application {
	return Application{
		AuthApplication: app.NewAuthService(service.UserService, service.TokenProvider, service.PasswordSecurity),
	}

}
