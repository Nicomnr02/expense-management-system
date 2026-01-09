package authservice

import (
	authdto "expense-management-system/cmd/auth/dto"
	authrepository "expense-management-system/cmd/auth/repository"
	"expense-management-system/pkg/httpserver"
	"expense-management-system/pkg/jwt"
	"expense-management-system/pkg/validator"
)

type AuthService interface {
	Login(ctx httpserver.Context, req authdto.LoginReq) (authdto.LoginRes, error)
}

type authserviceImpl struct {
	authrepository authrepository.Authrepository
	validator      validator.Validator
	JWTManager     *jwt.JWTManager
}

func New(authrepository authrepository.Authrepository,
	validator validator.Validator,
	JWTManager *jwt.JWTManager,
) AuthService {
	return &authserviceImpl{
		authrepository,
		validator,
		JWTManager,
	}
}
