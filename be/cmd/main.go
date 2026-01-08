package main

import (
	"context"

	"expense-management-system/cmd/auth"
	"expense-management-system/cmd/health"
	"expense-management-system/config"
	"expense-management-system/database"
	"expense-management-system/pkg/httpserver"
	"expense-management-system/pkg/jwt"
	"expense-management-system/pkg/logger"
	"expense-management-system/pkg/validator"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("starting application...")

	log.Println("loading configuration...")
	cfg := config.LoadConfig()

	log.Println("initializing logger...")
	logger := logger.InitZap(cfg.AppEnv)
	defer logger.Sync()

	log.Println("initializing database connection...")
	database := database.New(cfg)
	database.Init()

	log.Println("initializing server connection...")
	server := httpserver.New(httpserver.ServerConfig{
		ReadTimeout:  time.Duration(cfg.AppReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.AppWriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.AppIdleTimeout) * time.Second,
	})

	log.Println("initializing validator...")
	validator := validator.New()

	log.Println("initializing JWT Manager...")
	JWTManager := jwt.NewJWTManager(
		cfg.JWTAccessSecret,
		cfg.JWTRefreshSecret,
		cfg.AppName,
		cfg.JWTAccessTokenDurationMinutes,
		cfg.JWTRefreshTokenDurationMinutes,
	)

	log.Println("preparing modules...")
	health.Init(server, database)
	auth.Init(server, database, validator, JWTManager, logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		log.Printf("starting server on port %s...", cfg.AppPort)
		if err := server.App.Listen(":" + cfg.AppPort); err != nil {
			cancel()
		}
	}()

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		cancel()
	}()

	<-ctx.Done()

	log.Printf("Server shutdown initiated. Context: %v | Time: %s", ctx.Err(), time.Now().Format(time.RFC3339))

	log.Println("closing database pool...")
	database.Close()

	log.Println("closing scheduler...")

	log.Println("closing app...")
	server.Shutdown()

	log.Printf("Success shutdown gracefully. Context: %v | Time: %s", ctx.Err(), time.Now().Format(time.RFC3339))
}
