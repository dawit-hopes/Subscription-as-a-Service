package bootstrap

import (
	"github.com/dawit_hopes/saas/auth/internal/domain/port/outbound"
	"github.com/dawit_hopes/saas/auth/internal/infra/db"
		"go.mongodb.org/mongo-driver/mongo"

)

type Persistance struct {
	UserRepository outbound.UserRepository
}

func InitPersistence(collection *mongo.Collection) Persistance {
	return Persistance{
		UserRepository: db.NewUserRepositoryMongo(collection),
	}
}
