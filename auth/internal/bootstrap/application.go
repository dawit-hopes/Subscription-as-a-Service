package bootstrap

import (
	"github.com/dawit_hopes/saas/auth/internal/app"
	auth "github.com/dawit_hopes/saas/auth/internal/domain/port/inbound/auth"
)

type Application struct {
	AuthApplication auth.AuthService
}

func InitApplication(service Service) Application {
	return Application{
		AuthApplication: app.NewAuthService(service.UserService, service.TokenProvider, service.PasswordSecurity, service.TokenService),
	}

}
