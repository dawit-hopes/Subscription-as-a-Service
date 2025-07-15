package bootstrap

import (
	"github.com/dawit_hopes/saas/auth/internal/domain/port/outbound"
	"github.com/dawit_hopes/saas/auth/internal/infra/db/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Persistance struct {
	UserRepository outbound.UserRepository
	TokenRepositoy outbound.RefreshTokenRepository
}

func InitPersistence(collection []*mongo.Collection) Persistance {
	return Persistance{
		UserRepository: db.NewUserRepositoryMongo(collection[0]),
		TokenRepositoy: db.NewTokenRepository(collection[1]),
	}
}
