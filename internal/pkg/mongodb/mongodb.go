package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/lucasd-coder/backend-gobarber-golang/config"
	"github.com/lucasd-coder/backend-gobarber-golang/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func SetUpMongoDB(cfg *config.Config) {
	url := fmt.Sprintf("mongodb://%s:%d", cfg.MongoDbHost, cfg.MongoDbPort)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username: cfg.MongoDbUsername,
		Password: cfg.MongoDbPassword,
	}).ApplyURI(url))
	if err != nil {
		logger.Log.Fatal(err.Error())
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		logger.Log.Fatalf("Error MongoDB connection: %v", err.Error())
	} else {
		logger.Log.Infoln("MongoDB Connected ")
	}

	client = mongoClient
}

func GetClientMongoDB() *mongo.Client {
	return client
}

func CloseConnMongoDB() error {
	err := client.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	return nil
}
