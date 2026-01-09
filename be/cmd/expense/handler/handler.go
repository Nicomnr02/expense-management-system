package expensehandler

import (
	expenseservice "expense-management-system/cmd/expense/service"
	"expense-management-system/internal/middleware"
	"expense-management-system/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

type expenseHandlerImpl struct {
	server         *fiber.App
	expenseservice expenseservice.ExpenseService
	JWTManager     *jwt.JWTManager
}

func New(server *fiber.App, expenseservice expenseservice.ExpenseService,
	JWTManager *jwt.JWTManager) {
	handler := expenseHandlerImpl{
		server:         server,
		expenseservice: expenseservice,
		JWTManager:     JWTManager,
	}

	expenses := server.Group("expenses", middleware.Authenticate(JWTManager))

	expenses.Post("/", handler.CreateExpense)
	expenses.Get("/", handler.FetchExpense)
}
