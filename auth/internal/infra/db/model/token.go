package db

import (
	"time"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
	"github.com/dawit_hopes/saas/auth/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefreshTokenDocument struct {
	UserID    primitive.ObjectID `bson:"user_id"`
	Token     string             `bson:"token"`
	Revoked   bool               `bson:"revoked"`
	ExpiresAt time.Time          `bson:"expires_at"`
}

func (d RefreshTokenDocument) ToTokenModel() model.RefreshToken {
	return model.RefreshToken{
		UserID:    d.UserID.Hex(),
		Token:     d.Token,
		Revoked:   d.Revoked,
		ExpiresAt: d.ExpiresAt,
	}
}

func ToTokenDocument(token model.RefreshToken) (RefreshTokenDocument, *appErr.AppError) {
	objectID, err := primitive.ObjectIDFromHex(token.UserID)
	if err != nil {
		return RefreshTokenDocument{}, appErr.ErrInvalidID
	}

	return RefreshTokenDocument{
		UserID:    objectID,
		Token:     token.Token,
		Revoked:   token.Revoked,
		ExpiresAt: token.ExpiresAt,
	}, nil
}
