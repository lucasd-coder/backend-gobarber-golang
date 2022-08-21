package logger

import (
	"os"

	"github.com/lucasd-coder/backend-gobarber-golang/config"

	log "github.com/sirupsen/logrus"
)

var Log = log.WithFields(log.Fields{
	"logName":  "gobarber-golang",
	"logIndex": "message",
})

func SetUpLog(cfg *config.Config) {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	logLevel, _ := log.ParseLevel(cfg.Level)
	log.SetLevel(logLevel)
}
