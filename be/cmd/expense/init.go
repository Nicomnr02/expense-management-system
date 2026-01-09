package expense

import (
	authrepository "expense-management-system/cmd/auth/repository"
	expensehandler "expense-management-system/cmd/expense/handler"
	expenserepository "expense-management-system/cmd/expense/repository"
	expenseservice "expense-management-system/cmd/expense/service"
	"expense-management-system/config"
	"expense-management-system/database"
	"expense-management-system/pkg/jwt"
	"expense-management-system/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

func Init(
	server *fiber.App,
	database *database.Database,
	validator *validator.Validator,
	transaction *database.Transaction,
	JWTManager *jwt.JWTManager,
	config *config.Config,
) {
	authrepository := authrepository.New(database)
	expenseRepository := expenserepository.New(database)
	expenseService := expenseservice.New(
		authrepository,
		expenseRepository,
		transaction,
		validator,
		config,
	)

	expensehandler.New(server, expenseService, JWTManager)
}
