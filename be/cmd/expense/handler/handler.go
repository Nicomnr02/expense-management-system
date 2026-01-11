package expensehandler

import (
	expenseservice "expense-management-system/cmd/expense/service"
	"expense-management-system/internal/middleware"
	middlewareenum "expense-management-system/internal/middleware/enum"
	"expense-management-system/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

type expenseHandlerImpl struct {
	router         fiber.Router
	expenseservice expenseservice.ExpenseService
	JWTManager     *jwt.JWTManager
}

func New(router fiber.Router, expenseservice expenseservice.ExpenseService,
	JWTManager *jwt.JWTManager) {
	handler := expenseHandlerImpl{
		router:         router,
		expenseservice: expenseservice,
		JWTManager:     JWTManager,
	}

	expenses := router.Group("expenses", middleware.Authenticate(JWTManager))

	expenses.Post("/", handler.CreateExpense)
	expenses.Get("/", handler.FetchExpense)
	expenses.Get("/:id", handler.FetchExpenseDetail)
	expenses.Put("/:id/approve", middleware.AuthorizeRole(middlewareenum.MANAGERROLE), handler.ApproveExpense)
	expenses.Put("/:id/reject", middleware.AuthorizeRole(middlewareenum.MANAGERROLE), handler.RejectExpense)
}
