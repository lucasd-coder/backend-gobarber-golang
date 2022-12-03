//go:build integration
// +build integration

package repository

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lucasd-coder/backend-gobarber-golang/config"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	env_value         = "NotificationsRepository"
	testDocumentation = "documentation_examples"
	testCollection    = "notifications_test"
	testId            = "630b7bcb419f837457644cbc"
)

func factoryNotification() *model.Notification {
	id := primitive.NewObjectID()
	return &model.Notification{
		ID:          id,
		Content:     "Teste insert",
		RecipientID: uuid.MustParse("bd9a314b-c79e-4169-9594-ceef41964150"),
		Read:        true,
		CreatedAt:   time.Now().Local(),
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	if testing.Short() {
		log.Print("skipping mtest integration test in short mode")
		return
	}

	if err := mtest.Setup(); err != nil {
		log.Fatal(err)
	}
	defer os.Exit(m.Run())
	if err := mtest.Teardown(); err != nil {
		log.Fatal(err)
	}
}

func TestSaveData(t *testing.T) {
	cfg := SetUpConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mtest.ClusterURI()))

	defer client.Disconnect(ctx)

	collection := client.Database(testDocumentation).Collection(testCollection)

	cleanData(collection)

	require.NoError(t, err)

	repo := NotificationsRepository{cfg, client}

	notification := factoryNotification()

	err = repo.Save(notification)
	assert.Nil(t, err)
}

func SetUpConfig() *config.Config {
	err := setEnvValues()
	if err != nil {
		panic(err)
	}
	var cfg config.Config
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}

func insertData(collation *mongo.Collection) {
	_, err := collation.InsertOne(context.TODO(), factoryNotification())
	if err != nil {
		panic(err)
	}
}

func cleanData(collation *mongo.Collection) {
	_, err := collation.DeleteMany(context.TODO(), options.Delete())
	if err != nil {
		panic(err)
	}
}

func setEnvValues() error {
	err := os.Setenv("USERNAME_DB", env_value)
	if err != nil {
		return fmt.Errorf("Error setting USERNAME_DB, err = %v", err)
	}

	err = os.Setenv("PASSWORD_DB", env_value)
	if err != nil {
		return fmt.Errorf("Error setting PASSWORD_DB, err = %v", err)
	}

	err = os.Setenv("JWT_SECRET", env_value)
	if err != nil {
		return fmt.Errorf("Error setting JWT_SECRET, err = %v", err)
	}

	err = os.Setenv("JWT_ISSUER", env_value)
	if err != nil {
		return fmt.Errorf("Error setting JWT_ISSUER, err = %v", err)
	}

	err = os.Setenv("HOST_ETHEREAL_MAIL", env_value)
	if err != nil {
		return fmt.Errorf("Error setting HOST_ETHEREAL_MAIL, err = %v", err)
	}

	err = os.Setenv("PORT_ETHEREAL_MAIL", env_value)
	if err != nil {
		return fmt.Errorf("Error setting PORT_ETHEREAL_MAIL, err = %v", err)
	}

	err = os.Setenv("USERNAME_ETHEREAL_MAIL", env_value)
	if err != nil {
		return fmt.Errorf("Error setting USERNAME_ETHEREAL_MAIL, err = %v", err)
	}

	err = os.Setenv("PASSWORD_ETHEREAL_MAIL", env_value)
	if err != nil {
		return fmt.Errorf("Error setting PASSWORD_ETHEREAL_MAIL, err = %v", err)
	}

	err = os.Setenv("APP_NAME", env_value)
	if err != nil {
		return fmt.Errorf("Error setting APP_NAME, err = %v", err)
	}

	err = os.Setenv("APP_VERSION", env_value)
	if err != nil {
		return fmt.Errorf("Error setting APP_VERSION, err = %v", err)
	}

	err = os.Setenv("HTTP_PORT", "8080")
	if err != nil {
		return fmt.Errorf("Error setting HTTP_PORT, err = %v", err)
	}

	err = os.Setenv("LOG_LEVEL", "info")
	if err != nil {
		return fmt.Errorf("Error setting LOG_LEVEL, err = %v", err)
	}

	err = os.Setenv("APP_WEB_URL", "http://localhost:8080")
	if err != nil {
		return fmt.Errorf("Error setting APP_WEB_URL, err = %v", err)
	}

	err = os.Setenv("HOST_DB", "http://localhost:8080")
	if err != nil {
		return fmt.Errorf("Error setting HOST_DB, err = %v", err)
	}

	err = os.Setenv("HOST_MONGODB", "http://localhost:8080")
	if err != nil {
		return fmt.Errorf("Error setting HOST_MONGODB, err = %v", err)
	}

	err = os.Setenv("DATABASE_MONGODB", "http://localhost:8080")
	if err != nil {
		return fmt.Errorf("Error setting DATABASE_MONGODB, err = %v", err)
	}

	err = os.Setenv("PORT_MONGODB", "8080")
	if err != nil {
		return fmt.Errorf("Error setting PORT_MONGODB, err = %v", err)
	}

	err = os.Setenv("PORT_DB", "8080")
	if err != nil {
		return fmt.Errorf("Error setting PORT_DB, err = %v", err)
	}

	err = os.Setenv("PORT_MONGODB", "8080")
	if err != nil {
		return fmt.Errorf("Error setting PORT_MONGODB, err = %v", err)
	}

	err = os.Setenv("REDIS_DB", "3")
	if err != nil {
		return fmt.Errorf("Error setting REDIS_DB, err = %v", err)
	}

	err = os.Setenv("REDIS_PORT", "8080")
	if err != nil {
		return fmt.Errorf("Error setting REDIS_PORT, err = %v", err)
	}

	return nil
}
