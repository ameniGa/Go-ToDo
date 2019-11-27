package db

import (
	"context"
	utils "github.com/3almadmoon/ameni-assignment/api"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx context.Context

func Connect() (*mongo.Collection, error) {
	var err error
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(utils.DB_URI))
	if err != nil {
		return nil, err
	}
	log.Println("successfully connected to MongoDB")
	collection := client.Database(utils.DB_NAME).Collection(utils.DB_COLLECTION_NAME)
	return collection, err
}

func Disconnect() {
	client.Disconnect(ctx)
}
