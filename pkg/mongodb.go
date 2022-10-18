package pkg

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateIndex - creates an index for a specific field in a collection
// Creating indexes in MongoDB is an idempotent operation, so no need to check if it exists:
// https://stackoverflow.com/a/35020346/1794478
func CreateIndex(collection *mongo.Collection, field string, unique bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if createdIndex, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		// index in ascending order or -1 for descending order
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}); err != nil {
		return err
	} else {
		log.Printf("created index %s for field %s", createdIndex, field)
	}

	return nil
}
