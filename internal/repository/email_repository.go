package repository

import (
	"context"

	"github.com/venture-technology/venture/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type IEmailRepository interface {
	Record(ctx context.Context, email *entity.Email) error
}

type EmailRepository struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

func NewEmailRepository(client *mongo.Client, dbName, collectionName string, logger *zap.Logger) *EmailRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &EmailRepository{
		collection: collection,
		logger:     logger,
	}
}

func (er *EmailRepository) Record(ctx context.Context, email *entity.Email) error {
	_, err := er.collection.InsertOne(ctx, email)
	return err
}
