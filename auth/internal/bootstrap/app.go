package bootstrap

import (
	"github.com/dawit_hopes/saas/auth/internal/infra/config"
	"github.com/dawit_hopes/saas/auth/internal/infra/log"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

func InitializeApp() (*gin.Engine, error) {
	// Logger
	if err := log.InitLogger(); err != nil {
		return nil, err
	}
	defer log.Sync()

	env := config.NewEnv()

	mongoClient, err := config.NewMongoClient(env)
	if err != nil {
		log.Logger.Error("Failed to initialize MongoDB client: " + err.Error())
		return nil, err
	}

	db := mongoClient.Database(env.MongoDBName)

	persistance := InitPersistence([]*mongo.Collection{
		db.Collection("users"),
		db.Collection("tokens"),
	})
	signingKey := []byte(env.SigningKey)
	issuer := env.Issuer
	cost := env.Cost
	service := InitServices(signingKey, issuer, cost, persistance)
	application := InitApplication(service)

	// Router (with middleware + routes)
	router := InitRouter(application)

	return router, nil
}
