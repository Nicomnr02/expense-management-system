package auth

import (
	authhandler "expense-management-system/cmd/auth/handler"
	authrepository "expense-management-system/cmd/auth/repository"
	authservice "expense-management-system/cmd/auth/service"
	"expense-management-system/database"
	"expense-management-system/pkg/httpserver"
)

func Init(server httpserver.Server, database *database.Database) {
	authRepository := authrepository.New(database)
	authService := authservice.New(authRepository)
	authhandler.New(server, authService)
}

