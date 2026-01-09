package health

import (
	healthhandler "expense-management-system/cmd/health/handler"
	"expense-management-system/database"
	"expense-management-system/pkg/httpserver"
)

func Init(server *httpserver.Server, database *database.Database) {
	healthhandler.New(server, database)
}
