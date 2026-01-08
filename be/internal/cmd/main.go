package main

import (
	"context"
	"expense-management-system/internal/config"
	"expense-management-system/internal/database"
	"expense-management-system/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("starting application...")

	log.Println("loading configuration...")
	cfg := config.LoadConfig()

	log.Println("initializing logger...")
	logger := logger.InitZap(cfg.AppEnv)
	defer logger.Sync()

	log.Println("initializing database connection...")
	database := database.Init(cfg)

	log.Println("initializing server connection...")
	app := fiber.New(
		fiber.Config{
			ReadTimeout:  time.Duration(cfg.AppReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.AppWriteTimeout) * time.Second,
			IdleTimeout:  time.Duration(cfg.AppIdleTimeout) * time.Second,
		},
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		log.Printf("starting server on port %s...", cfg.AppPort)
		if err := app.Listen(":" + cfg.AppPort); err != nil {
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
	_ = app.Shutdown()

	log.Printf("Success shutdown gracefully. Context: %v | Time: %s", ctx.Err(), time.Now().Format(time.RFC3339))
}
