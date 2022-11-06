package database

import (
	"fmt"
	"time"

	"github.com/lucasd-coder/backend-gobarber-golang/config"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/pkg/migrations"
	"github.com/lucasd-coder/backend-gobarber-golang/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

var db *gorm.DB

func StartDB(cfg *config.Config) {
	str := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s search_path=%s",
		cfg.Host, cfg.PostgresPort, cfg.Username, cfg.Dbname, cfg.Password, cfg.Schema)

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		logger.Log.Fatal(err.Error())
	} else {
		logger.Log.Infoln("Postgres Connected")
	}

	database.Use(prometheus.New(
		prometheus.Config{
			DBName:          cfg.Dbname,
			RefreshInterval: 15,
			StartServer:     false,
			MetricsCollector: []prometheus.MetricsCollector{
				&prometheus.Postgres{
					VariableNames: []string{"Threads_running"},
				},
			},
		},
	))

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
