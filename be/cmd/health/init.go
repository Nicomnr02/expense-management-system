package health

import (
	healthhandler "expense-management-system/cmd/health/handler"
	"expense-management-system/database"

	"github.com/gofiber/fiber/v2"
)

func Init(server *fiber.App, database *database.Database) {
	healthhandler.New(server, database)
}
