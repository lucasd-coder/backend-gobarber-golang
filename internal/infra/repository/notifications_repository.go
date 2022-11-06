package repository

import (
	"context"

	"github.com/lucasd-coder/backend-gobarber-golang/config"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationsRepository struct {
	Config     *config.Config
	Connection *mongo.Client
}

func NewNotificationsRepository(cfg *config.Config, connectionDb *mongo.Client) *NotificationsRepository {
	return &NotificationsRepository{
		Config:     cfg,
		Connection: connectionDb,
	}
}

func (repo *NotificationsRepository) Save(notification *model.Notification) error {
	collection := repo.Connection.Database(repo.Config.MongoDbDatabase).Collection("notifications")

	_, err := collection.InsertOne(context.TODO(), notification)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}
