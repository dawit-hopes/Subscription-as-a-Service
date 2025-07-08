package bootstrap

import (
	"github.com/dawit_hopes/saas/auth/internal/infra/log"
	"github.com/gin-gonic/gin"
)

func InitializeApp() (*gin.Engine, error) {
	// Logger
	if err := log.InitLogger(); err != nil {
		return nil, err
	}
	defer log.Sync()

	// Router (with middleware + routes)
	router := IninRouter()

	return router, nil
}
