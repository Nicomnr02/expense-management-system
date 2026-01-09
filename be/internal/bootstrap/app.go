package bootstrap

import (
	"context"
	"expense-management-system/cmd/auth"
	"expense-management-system/cmd/expense"
	"expense-management-system/cmd/health"
	"expense-management-system/config"
	"expense-management-system/database"
	"expense-management-system/internal/middleware"
	"expense-management-system/pkg/jwt"
	"expense-management-system/pkg/logger"
	"expense-management-system/pkg/validator"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Run() error {
	log := logger.InitZap()
	defer log.Sync()

	log.Info("starting application...")

	log.Info("loading configuration...")
	cfg := config.LoadConfig()

	log.Info("loading timezone...")
	_, err := time.LoadLocation(cfg.AppTimezone)
	if err != nil {
		log.Error("Failed to load timezone",
			zap.String("tz", cfg.AppTimezone),
			zap.Error(err),
		)
		return err
	}

	log.Info("initializing database connection...")
	db := database.New(cfg)
	if err := db.Init(); err != nil {
		log.Error(err.Error())
		return err
	}

	log.Info("initializing database transaction...")
	tx := database.NewTransaction(db)

	log.Info("initializing server connection...")
	server := fiber.New(fiber.Config{
		ReadTimeout:  time.Duration(cfg.AppReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.AppWriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.AppIdleTimeout) * time.Second,
	})

	log.Info("initializing JWT Manager...")
	JWTManager := jwt.NewJWTManager(
		cfg.JWTAccessSecret,
		cfg.JWTRefreshSecret,
		cfg.AppName,
		cfg.JWTAccessTokenDurationMinutes,
		cfg.JWTRefreshTokenDurationMinutes,
	)

	log.Info("register middlewares...")
	server.Use(
		middleware.RequestID(log),
		middleware.RequestLogger(),
	)

	log.Info("initializing validator...")
	validator := validator.New()

	log.Info("register modules...")
	health.Init(server, db)
	auth.Init(server, db, validator, JWTManager)
	expense.Init(server, db, validator, tx, JWTManager, cfg)

	RunServer(server, db, cfg, log)

	return nil
}

func RunServer(server *fiber.App, database *database.Database, cfg *config.Config, logger *zap.Logger) {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		logger.Info("starting server on port " + cfg.AppPort)
		if err := server.Listen(":" + cfg.AppPort); err != nil {
			logger.Error(err.Error())
			stop()
		}
	}()

	<-ctx.Done()

	ShutdownServer(ctx, server, database, logger)
}

func ShutdownServer(ctx context.Context, server *fiber.App, database *database.Database, logger *zap.Logger) {
	logger.Info("shutdown initiated", zap.String("reason", ctx.Err().Error()))

	logger.Info("closing database pool...")
	database.Close()

	logger.Info("closing scheduler...")

	logger.Info("closing app...")
	server.Shutdown()

	logger.Info("shutdown completed")
}
