package logger

import (
	"expense-management-system/config"
	"log"

	"go.uber.org/zap"
)

func InitZap(cfg *config.Config) *zap.Logger {
	var (
		zlog *zap.Logger
		err  error
	)
	if cfg.AppEnv == "production" {
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
