package logger

import (
	"log"

	"go.uber.org/zap"
)

func InitZap() *zap.Logger {
	cfg := zap.NewProductionConfig()

	cfg.DisableStacktrace = true

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	return logger
}
