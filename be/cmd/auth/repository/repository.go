package authrepository

import (
	"context"
	authdomain "expense-management-system/cmd/auth/domain"
	authquery "expense-management-system/cmd/auth/repository/query"
	"expense-management-system/database"
)

type Authrepository interface {
	FetchUser(c context.Context, q authquery.FetchUser) (authdomain.User, error)
}
type AuthrepositoryImpl struct {
	DB *database.Database
}

func New(DB *database.Database) Authrepository {
	return &AuthrepositoryImpl{DB}
}
