package database

import (
	"fmt"
	"time"

	"backend-gobarber-golang/config"
	"backend-gobarber-golang/internal/pkg/migrations"
	"backend-gobarber-golang/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB(cfg *config.Config) {
	str := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s search_path=%s",
		cfg.Host, cfg.PostgresPort, cfg.Username, cfg.Dbname, cfg.Password, cfg.Schema)

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal(err.Error())
	} else {
		logger.Log.Infoln("Connected")
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(cfg.MaxIdleConns)
	config.SetMaxOpenConns(cfg.MaxOpenConns)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
