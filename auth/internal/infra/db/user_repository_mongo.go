package db

import (
	"context"

	"github.com/dawit_hopes/saas/auth/internal/domain/model"
	"github.com/dawit_hopes/saas/auth/internal/domain/port/outbound"
	"github.com/dawit_hopes/saas/auth/internal/infra/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
)

type userRepositoryMongo struct {
	collection *mongo.Collection
}

func NewUserRepositoryMongo(collection *mongo.Collection) outbound.UserRepository {
	return &userRepositoryMongo{
		collection: collection,
	}
}
func (userRepo *userRepositoryMongo) GetByID(ctx context.Context, id string) (*model.User, *appErr.AppError) {
	filter := bson.M{"_id": id}
	var user UserDocument
	err := userRepo.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, appErr.ErrGeneralDatabaseQuery
	}
	u := user.ToDomain()
	return &u, nil
}

func (userRepo *userRepositoryMongo) Create(ctx context.Context, user *model.User) (*model.User, *appErr.AppError) {
	userDoc, err := FromDomain(*user)
	if err != nil {
		log.Logger.Error("failed to convert user to document: " + err.Error())
		return nil, appErr.ErrInternalServer
	}

	_, err = userRepo.collection.InsertOne(ctx, userDoc)
	if err != nil {
		log.Logger.Error("failed to insert user into database: " + err.Error())
		return nil, appErr.ErrGeneralDatabaseInsert
	}

	user.CreatedAt = userDoc.CreatedAt.Time()
	user.UpdatedAt = userDoc.UpdatedAt.Time()
	user.ID = userDoc.ID.Hex()

	return user, nil
}

func (userRepo *userRepositoryMongo) Update(ctx context.Context, user *model.User) (*model.User, *appErr.AppError) {
	userDoc, err := FromDomain(*user)
	if err != nil {
		log.Logger.Error("failed to convert user to document: " + err.Error())
		return nil, appErr.ErrGeneralDatabaseUpdate
	}

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": userDoc}

	_, err = userRepo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Logger.Error("failed to update user in database: " + err.Error())
		return nil, appErr.ErrGeneralDatabaseUpdate
	}

	user.ID = userDoc.ID.Hex()
	user.UpdatedAt = userDoc.UpdatedAt.Time()

	return user, nil
}

func (userRepo *userRepositoryMongo) Delete(ctx context.Context, id string) *appErr.AppError {
	objecID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Logger.Error("failed to convert id to ObjectID: " + err.Error())
		return appErr.ErrGeneralDatabaseDelete
	}
	filter := bson.M{"_id": objecID}
	_, err = userRepo.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Logger.Error("failed to delete user from database: " + err.Error())
		return appErr.ErrGeneralDatabaseDelete
	}
	return nil
}

func (userRepo *userRepositoryMongo) GetByEmail(ctx context.Context, email string) (*model.User, *appErr.AppError) {
	filter := bson.M{"email": email}
	var user UserDocument
	err := userRepo.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, appErr.ErrGeneralDatabaseQuery
	}
	userDomain := user.ToDomain()
	return &userDomain, nil
}
