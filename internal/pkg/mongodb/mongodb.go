package mongodb

import (
	"context"
	"fmt"
	"time"

	"backend-gobarber-golang/config"
	"backend-gobarber-golang/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func SetUpMongoDB(cfg *config.Config) {
	url := fmt.Sprintf("mongodb://%s:%d", cfg.MongoDbHost, cfg.MongoDbPort)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		logger.Log.Fatal(err.Error())
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		logger.Log.Fatalf("Error MongoDB connection: ", err.Error())
	} else {
		logger.Log.Infoln("MongoDB Connected ")
	}

	client = mongoClient

	defer client.Disconnect(ctx)
}

func GetClientMongoDB() *mongo.Client {
	return client
}
