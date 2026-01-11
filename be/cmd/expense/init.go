package expense

import (
	authrepository "expense-management-system/cmd/auth/repository"
	expenseenum "expense-management-system/cmd/expense/enum"
	expensehandler "expense-management-system/cmd/expense/handler"
	expenserepository "expense-management-system/cmd/expense/repository"
	expenseservice "expense-management-system/cmd/expense/service"
	"expense-management-system/config"
	"expense-management-system/database"
	"expense-management-system/internal/job"
	"expense-management-system/pkg/jwt"
	"expense-management-system/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

func Init(
	server *fiber.App,
	database *database.Database,
	jobClient *job.Client,
	jobServer *job.Server,
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
		jobClient,
	)

	expensehandler.New(server, expenseService, JWTManager)

	jobServer.RegisterWorker(expenseenum.Pay, expenseService.PayExpense)
}
