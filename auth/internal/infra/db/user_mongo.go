// Package db provides MongoDB implementations for user persistence and conversion utilities.
package db

import (
	"github.com/dawit_hopes/saas/auth/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDocument struct {
	ID        primitive.ObjectID `bson:"_id"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

// ToDomain converts UserDocument to domain model User
func (u *UserDocument) ToDomain() model.User {
	return model.User{
		ID:        u.ID.Hex(),
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt.Time(),
		UpdatedAt: u.UpdatedAt.Time(),
	}
}

// FromDomain converts domain model User to UserDocument
func FromDomain(user model.User) (UserDocument, error) {
	id := primitive.NewObjectID()
	return UserDocument{
		ID:        id,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: primitive.NewDateTimeFromTime(user.CreatedAt),
		UpdatedAt: primitive.NewDateTimeFromTime(user.UpdatedAt),
	}, nil
}
