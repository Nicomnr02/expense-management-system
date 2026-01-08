package logger

import (
	"log"

	"go.uber.org/zap"
)

func InitZap(ENV string) *zap.Logger {
	var (
		zlog *zap.Logger
		err  error
	)
	if ENV == "production" {
		zlog, err = zap.NewProduction()
	} else {
		zlog, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatalf("Failed to init logger: %s", err.Error())
	}
	defer zlog.Sync()
	return zlog
}
