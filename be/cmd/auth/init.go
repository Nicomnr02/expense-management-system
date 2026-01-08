package auth

import (
	authhandler "expense-management-system/cmd/auth/handler"
	authrepository "expense-management-system/cmd/auth/repository"
	authservice "expense-management-system/cmd/auth/service"
	"expense-management-system/database"
	"expense-management-system/pkg/httpserver"
	"expense-management-system/pkg/jwt"
	"expense-management-system/pkg/validator"

	"go.uber.org/zap"
)

func Init(
	server httpserver.Server,
	database *database.Database,
	validator validator.Validator,
	JWTManager *jwt.JWTManager,
	logger *zap.Logger,
) {
	authRepository := authrepository.New(database)
	authService := authservice.New(authRepository, validator, JWTManager, logger)
	authhandler.New(server, authService, logger)
}
