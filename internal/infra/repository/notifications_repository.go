package repository

import (
	"context"

	"backend-gobarber-golang/internal/model"
	"backend-gobarber-golang/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationsRepository struct {
	Connection *mongo.Client
}

func NewNotificationsRepository(connectionDb *mongo.Client) *NotificationsRepository {
	return &NotificationsRepository{
		Connection: connectionDb,
	}
}

func (repo *NotificationsRepository) Save(notification *model.Notification) error {
	collection := repo.Connection.Database("gobarber").Collection("notifications")

	_, err := collection.InsertOne(context.TODO(), notification)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}
