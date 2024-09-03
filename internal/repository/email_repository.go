package repository

import (
	"context"

	"github.com/venture-technology/venture/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type IEmailRepository interface {
	Record(ctx context.Context, email *entity.Email) error
}

type EmailRepository struct {
	collection *mongo.Collection
}

func NewEmailRepository(client *mongo.Client, dbName, collectionName string) *EmailRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &EmailRepository{
		collection: collection,
	}
}

func (er *EmailRepository) Record(ctx context.Context, email *entity.Email) error {
	_, err := er.collection.InsertOne(ctx, email)
	return err
}
