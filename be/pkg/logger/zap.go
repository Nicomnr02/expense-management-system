package logger

import (
	"log"

	"go.uber.org/zap"
)

func InitZap(ENV string) *zap.Logger {
	zlog, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to init logger: %s", err.Error())
	}
	return zlog
}
