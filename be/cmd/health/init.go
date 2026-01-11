package health

import (
	healthhandler "expense-management-system/cmd/health/handler"
	"expense-management-system/database"
	"expense-management-system/internal/job"

	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router, database *database.Database,
	jobServer job.Server) {
	healthhandler.New(router, database, jobServer)
}
