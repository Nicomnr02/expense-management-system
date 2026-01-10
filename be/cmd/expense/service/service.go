package expenseservice

import (
	authrepository "expense-management-system/cmd/auth/repository"
	expensedto "expense-management-system/cmd/expense/dto"
	expenserepository "expense-management-system/cmd/expense/repository"
	"expense-management-system/config"
	"expense-management-system/database"
	"expense-management-system/model"
	"expense-management-system/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

type ExpenseService interface {
	CreateExpense(c *fiber.Ctx, req expensedto.CreateExpenseReq) (expensedto.CreateExpenseRes, error)

	FetchExpense(c *fiber.Ctx, req expensedto.FetchExpenseReq) ([]expensedto.FetchExpenseRes, model.Pagination, error)
	FetchExpenseDetail(c *fiber.Ctx, req expensedto.FetchExpenseDetailReq) (expensedto.FetchExpenseDetailRes, error)
}

type expenseServiceImpl struct {
	authrepository    authrepository.Authrepository
	expenserepository expenserepository.ExpenseRepository
	transaction       *database.Transaction
	validator         *validator.Validator
	cfg               *config.Config
}

func New(
	authrepository authrepository.Authrepository,
	expenserepository expenserepository.ExpenseRepository,
	transaction *database.Transaction,
	validator *validator.Validator,
	cfg *config.Config,
) ExpenseService {
	return &expenseServiceImpl{
		authrepository,
		expenserepository,
		transaction,
		validator,
		cfg,
	}
}
