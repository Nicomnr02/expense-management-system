package bootstrap

import (
	"context"
	"expense-management-system/cmd/auth"
	"expense-management-system/cmd/expense"
	"expense-management-system/cmd/health"
	"expense-management-system/config"
	"expense-management-system/database"
	"expense-management-system/internal/job"
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

	log.Info("initializing database connection...")
	db := database.New(cfg)
	if err := db.Init(); err != nil {
		log.Error(err.Error())
		return err
	}

	log.Info("initializing database transaction...")
	tx := database.NewTransaction(db)

	log.Info("initializing job client connection...")
	jobClient := job.NewClient(cfg)

	log.Info("initializing job server connection...")
	jobServer := job.NewServer(cfg, log)

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

	log.Info("register loggers...")
	server.Use(
		middleware.RequestID(log),
		middleware.RequestLogger(),
	)

	log.Info("initializing validator...")
	validator := validator.New()

	log.Info("register modules...")
	health.Init(server, db, jobServer)
	auth.Init(server, db, validator, JWTManager)
	expense.Init(server, db, jobClient, jobServer, validator, tx, JWTManager, cfg)

	RunServer(server, db, jobClient, jobServer, cfg, log)

	return nil
}

func RunServer(
	server *fiber.App,
	database *database.Database,
	jobClient *job.Client,
	jobServer *job.Server,
	cfg *config.Config,
	log *zap.Logger,
) {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		log.Info("starting job server on " + cfg.RedisAddr)
		if err := jobServer.Run(); err != nil {
			log.Error(err.Error())
			stop()
		}
	}()

	go func() {
		log.Info("starting server on port " + cfg.AppPort)
		if err := server.Listen(":" + cfg.AppPort); err != nil {
			log.Error(err.Error())
			stop()
		}
	}()

	<-ctx.Done()

	ShutdownServer(ctx, server, database, jobClient, jobServer, log)
}

func ShutdownServer(
	ctx context.Context,
	server *fiber.App,
	database *database.Database,
	jobClient *job.Client,
	jobServer *job.Server,
	log *zap.Logger,
) {
	log.Info("shutdown initiated", zap.String("reason", ctx.Err().Error()))

	log.Info("closing database connection...")
	database.Close()

	log.Info("closing job client connection...")
	_ = jobClient.Close()

	log.Info("closing job server connection...")
	jobServer.Shutdown()

	log.Info("closing server connection...")
	_ = server.Shutdown()

	log.Info("shutdown completed")
}
