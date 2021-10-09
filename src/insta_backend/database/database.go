package database

import (
	"context"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Database(ctx context.Context) *mongo.Client{
	var uri string = os.Getenv("MONGO_URI")
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	
	return client
}