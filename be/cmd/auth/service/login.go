package authservice

import (
	"errors"
	authdto "expense-management-system/cmd/auth/dto"
	authquery "expense-management-system/cmd/auth/repository/query"
	"expense-management-system/internal/contextkey"
	"expense-management-system/model"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (s *authserviceImpl) Login(c *fiber.Ctx, req authdto.LoginReq) (authdto.LoginRes, error) {

	var (
		ctx  = c.Context()
		log  = c.Locals(contextkey.Logger).(*zap.Logger)
		data authdto.LoginRes
	)

	err := s.validator.Validate.Struct(&req)
	if err != nil {
		log.Warn(err.Error())
		return data, model.ErrBadRequest("Invalid email or password")
	}

	user, err := s.authrepository.FetchUser(ctx, authquery.FetchUser{
		Email: req.Email,
	})
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Error(err.Error())
			return data, model.ErrInternalServer("Authentication failed")
		}
		return data, model.ErrBadRequest("Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return data, model.ErrBadRequest("Invalid email or password")
	}

	accessToken, refreshToken, err := s.JWTManager.GenerateTokens(user.ID, user.Role, user.Email)
	if err != nil {
		log.Error(err.Error())
		return data, model.ErrInternalServer("Authentication failed")
	}

	data = authdto.LoginRes{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}

	return data, nil
}
