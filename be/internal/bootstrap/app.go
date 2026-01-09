package bootstrap

import (
	"context"
	"expense-management-system/cmd/auth"
	"expense-management-system/cmd/health"
	"expense-management-system/config"
	"expense-management-system/database"
	"expense-management-system/internal/middleware"
	"expense-management-system/pkg/httpserver"
	"expense-management-system/pkg/jwt"
	"expense-management-system/pkg/logger"
	"expense-management-system/pkg/validator"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func Run() error {
	logger := logger.InitZap()
	defer logger.Sync()

	logger.Info("starting application...")

	logger.Info("loading configuration...")
	cfg := config.LoadConfig()

	logger.Info("initializing database connection...")
	database := database.New(cfg)
	if err := database.Init(); err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("initializing server connection...")
	server := httpserver.New(httpserver.ServerConfig{
		ReadTimeout:  time.Duration(cfg.AppReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.AppWriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.AppIdleTimeout) * time.Second,
	})

	logger.Info("register middlewares...")
	server.Middleware(middleware.RequestID(logger))
	server.Middleware(middleware.RequestLogger())

	logger.Info("initializing validator...")
	validator := validator.New()

	logger.Info("initializing JWT Manager...")
	JWTManager := jwt.NewJWTManager(
		cfg.JWTAccessSecret,
		cfg.JWTRefreshSecret,
		cfg.AppName,
		cfg.JWTAccessTokenDurationMinutes,
		cfg.JWTRefreshTokenDurationMinutes,
	)

	logger.Info("register modules...")
	health.Init(server, database)
	auth.Init(server, database, validator, JWTManager)

	RunServer(server, database, cfg, logger)

	return nil
}

func RunServer(server *httpserver.Server, database *database.Database, cfg *config.Config, logger *zap.Logger) {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		logger.Info("starting server on port " + cfg.AppPort)
		if err := server.App.Listen(":" + cfg.AppPort); err != nil {
			logger.Error(err.Error())
			stop()
		}
	}()

	<-ctx.Done()

	ShutdownServer(ctx, server, database, logger)
}

func ShutdownServer(ctx context.Context, server *httpserver.Server, database *database.Database, logger *zap.Logger) {
	logger.Info("shutdown initiated", zap.String("reason", ctx.Err().Error()))

	logger.Info("closing database pool...")
	database.Close()

	logger.Info("closing scheduler...")

	logger.Info("closing app...")
	server.Shutdown()

	logger.Info("shutdown completed")
}
