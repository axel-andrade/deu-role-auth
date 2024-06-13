package mongo_repositories

import (
	"context"

	mongo_database "github.com/axel-andrade/deu-role-auth/internal/adapters/secondary/database/mongo"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseRepository struct {
	collection *mongo.Collection
}

func BuildBaseRepository(collectionName string) *BaseRepository {
	db := mongo_database.GetDB()
	collection := db.Collection(collectionName)

	return &BaseRepository{collection: collection}
}

func (r *BaseRepository) StartTransaction() error {
	// Not implemented. Maybe use uow pattern for transactions in mongodb
	return nil
}

func (r *BaseRepository) CommitTransaction() error {
	// Not implemented. Maybe use uow pattern for transactions in mongodb
	return nil
}

func (r *BaseRepository) CancelTransaction() error {
	// Not implemented. Maybe use uow pattern for transactions in mongodb
	return nil
}

func (r *BaseRepository) NextEntityID() string {
	return uuid.NewV4().String()
}

func (r *BaseRepository) Create(data any) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.Background(), data)
}

func (r *BaseRepository) FindOne(filter any) *mongo.SingleResult {
	return r.collection.FindOne(context.Background(), filter)
}
