package db

import (
	"context"
	"errors"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
	"github.com/dawit_hopes/saas/auth/internal/domain/model"
	"github.com/dawit_hopes/saas/auth/internal/domain/port/outbound"
	db "github.com/dawit_hopes/saas/auth/internal/infra/db/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type tokenRepository struct {
	collection *mongo.Collection
}

func NewTokenRepository(collection *mongo.Collection) outbound.RefreshTokenRepository {
	return &tokenRepository{
		collection: collection,
	}
}

func (t *tokenRepository) EnsureTTLIndex() error {
	indexModel := mongo.IndexModel{
		Keys: bson.M{"expires_at": 1},
		Options: options.Index().
			SetExpireAfterSeconds(0).
			SetName("expires_at_ttl"),
	}

	_, err := t.collection.Indexes().CreateOne(context.Background(), indexModel)
	return err
}

func (t *tokenRepository) getObjectID(id string) (primitive.ObjectID, *appErr.AppError) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return primitive.ObjectID{}, appErr.ErrInvalidUserID
	}

	return objectID, nil
}

func (t *tokenRepository) Create(ctx context.Context, token model.RefreshToken) *appErr.AppError {
	tokenDoc, err := db.ToTokenDocument(token)
	if err != nil {
		return err
	}
	if _, err := t.collection.InsertOne(ctx, tokenDoc); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return appErr.ErrTokenAlreadyExists
		}
		return appErr.ErrGeneralDatabaseInsert
	}
	return nil
}
func (t *tokenRepository) GetByUserID(ctx context.Context, userID string) (*model.RefreshToken, *appErr.AppError) {
	objID, err := t.getObjectID(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"user_id": objID}

	var tokenDoc db.RefreshTokenDocument
	if err := t.collection.FindOne(ctx, filter).Decode(&tokenDoc); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, appErr.ErrUserNotFound
		}
		return nil, appErr.ErrGeneralDatabaseQuery
	}

	token := tokenDoc.ToTokenModel()
	return &token, nil
}
func (t *tokenRepository) Delete(ctx context.Context, userID string) *appErr.AppError {
	objID, err := t.getObjectID(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"user_id": objID}

	if _, err := t.collection.DeleteOne(ctx, filter); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return appErr.ErrDocumentNotFound
		}
	}

	return nil
}
func (t *tokenRepository) Update(ctx context.Context, req model.RefreshToken) *appErr.AppError {
	objID, err := t.getObjectID(req.UserID)
	if err != nil {
		return err
	}

	tokenDoc, err := db.ToTokenDocument(req)
	if err != nil {
		return err
	}

	filter := bson.M{"user_id": objID}
	update := bson.M{"$set": tokenDoc}

	if _, err := t.collection.UpdateOne(ctx, filter, update); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return appErr.ErrDocumentNotFound
		}
	}
	return nil
}

func (t *tokenRepository) RevokeByUserID(ctx context.Context, userID string) *appErr.AppError {
	objID, err := t.getObjectID(userID)
	if err != nil {
		return err
	}
	filter := bson.M{"user_id": objID}
	update := bson.M{"$set": bson.M{"revoked": true}}
	if _, err := t.collection.UpdateOne(ctx, filter, update); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return appErr.ErrDocumentNotFound
		}
		return appErr.ErrFailedToRevokToken
	}
	
	return nil
}
