package auth

import (
	authhandler "expense-management-system/cmd/auth/handler"
	authrepository "expense-management-system/cmd/auth/repository"
	authservice "expense-management-system/cmd/auth/service"
	"expense-management-system/database"
	"expense-management-system/pkg/jwt"
	"expense-management-system/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

func Init(
	router fiber.Router,
	database *database.Database,
	validator validator.Validator,
	JWTManager *jwt.JWTManager,
) {
	authRepository := authrepository.New(database)
	authService := authservice.New(authRepository, validator, JWTManager)
	authhandler.New(router, authService)
}
