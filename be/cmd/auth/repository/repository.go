package authrepository

import (
	"expense-management-system/database"
)

type Authrepository interface {
}
type AuthrepositoryImpl struct {
	DB *database.Database
}

func New(DB *database.Database) Authrepository {
	return &AuthrepositoryImpl{DB}
}
