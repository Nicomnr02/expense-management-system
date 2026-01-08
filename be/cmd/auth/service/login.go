package authservice

import (
	"context"
	"errors"
	authdto "expense-management-system/cmd/auth/dto"
	authquery "expense-management-system/cmd/auth/repository/query"
	"expense-management-system/dto"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *authserviceImpl) Login(c context.Context, req authdto.LoginReq) (authdto.LoginRes, error) {
	var data authdto.LoginRes

	err := s.validator.Validate.Struct(&req)
	if err != nil {
		s.logger.Warn(err.Error())
		return data, dto.ErrBadRequest("Invalid email or password")
	}

	user, err := s.authrepository.FetchUser(c, authquery.FetchUser{
		Email: req.Email,
	})
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			s.logger.Error(err.Error())
			return data, dto.ErrInternalServer("Authentication failed")
		}
		return data, dto.ErrBadRequest("Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return data, dto.ErrBadRequest("Invalid email or password")
	}

	accessToken, refreshToken, err := s.JWTManager.GenerateTokens(user.ID, user.Role, user.Email)
	if err != nil {
		s.logger.Error(err.Error())
		return data, dto.ErrInternalServer("Authentication failed")
	}

	data = authdto.LoginRes{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}

	return data, nil
}
