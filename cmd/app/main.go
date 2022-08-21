package main

import (
	"os"

	"github.com/lucasd-coder/backend-gobarber-golang/config"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/app"
	"github.com/lucasd-coder/backend-gobarber-golang/pkg/logger"

	"github.com/ilyakaznacheev/cleanenv"
)

var cfg config.Config

func main() {
	// Profile: dev, homolog, prod
	profile := os.Getenv("GO_PROFILE")
	var path string

	switch profile {
	case "dev":
		path = "./config/config-dev.yml"
		break
	}

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		logger.Log.Fatalf("Config error: %v", err)
	}
	config.ExportConfig(&cfg)

	app.Run(&cfg)
}
