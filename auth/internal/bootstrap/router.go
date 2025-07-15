// Package bootstrap provides initialization routines for setting up the application's router.
package bootstrap

import (
	"github.com/dawit_hopes/saas/auth/internal/infra/http/middleware"
	"github.com/dawit_hopes/saas/auth/internal/infra/http/routes"

	"github.com/dawit_hopes/saas/auth/internal/infra/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var APIVersion = "v1"

func InitRouter(application Application, service Service) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHandler(log.Logger))

	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Logger.Fatal("failed to set trusted proxies", zap.Error(err))
	}

	v1 := r.Group("/api/" + APIVersion)
	protected := v1.Group("/auth")
	protected.Use(middleware.TokenMiddleware(service.TokenProvider))

	r.GET("/health", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, gin.H{"status": "ok"})
	})

	routes.RegisterAuthRoutes(v1.Group("/auth"), application.AuthApplication)
	routes.RegisterProtectedAuthRoutes(protected, application.AuthApplication)

	return r
}
