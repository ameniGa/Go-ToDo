package db

import (
	"context"
	"log"
	"time"

	config "github.com/3almadmoon/ameni-assignment/configs"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	ctx    context.Context
	cancel context.CancelFunc
)

//init function parse config file to get db params
func init() {
	if err := config.SetViper(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

//Connect to mongoDB , create database and collection if don't exist
func Connect() (*mongo.Collection, error) {
	var err error
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("database.uri")))
	if err != nil {
		return nil, err
	}
	log.Println("successfully connected to MongoDB")
	collection := client.Database(viper.GetString("database.name")).Collection(viper.GetString("database.collection"))
	return collection, err
}

//Disconnect from mongoDB
func Disconnect() {
	client.Disconnect(ctx)
}
