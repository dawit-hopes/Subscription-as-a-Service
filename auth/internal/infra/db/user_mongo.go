// Package db provides MongoDB implementations for user persistence and conversion utilities.
package db

import (
	"github.com/dawit_hopes/saas/auth/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserMongo struct {
	ID        primitive.ObjectID `bson:"_id"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

// ToDomain converts UserMongo to domain model User
func (u *UserMongo) ToDomain() model.User {
	return model.User{
		ID:        u.ID.Hex(),
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt.Time(),
		UpdatedAt: u.UpdatedAt.Time(),
	}
}

// FromDomain converts domain model User to UserMongo
func FromDomain(user model.User) (UserMongo, error) {
	id, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return UserMongo{}, err
	}
	return UserMongo{
		ID:        id,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: primitive.NewDateTimeFromTime(user.CreatedAt),
		UpdatedAt: primitive.NewDateTimeFromTime(user.UpdatedAt),
	}, nil
}
