package authservice

import (
	"context"
	authdto "expense-management-system/cmd/auth/dto"
	authrepository "expense-management-system/cmd/auth/repository"
	"expense-management-system/pkg/jwt"
	"expense-management-system/pkg/validator"

	"go.uber.org/zap"
)

type AuthService interface {
	Login(ctx context.Context, req authdto.LoginReq) (authdto.LoginRes, error)
}

type authserviceImpl struct {
	authrepository authrepository.Authrepository
	validator      validator.Validator
	JWTManager     *jwt.JWTManager
	logger         *zap.Logger
}

func New(authrepository authrepository.Authrepository,
	validator validator.Validator,
	JWTManager *jwt.JWTManager,
	logger *zap.Logger,
) AuthService {
	return &authserviceImpl{
		authrepository,
		validator,
		JWTManager,
		logger,
	}
}
