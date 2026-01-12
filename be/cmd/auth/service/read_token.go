package authservice

import (
	"errors"
	authdto "expense-management-system/cmd/auth/dto"
	authquery "expense-management-system/cmd/auth/repository/query"
	"expense-management-system/internal/contextkey"
	"expense-management-system/model"
	"expense-management-system/pkg/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (s *authserviceImpl) ReadToken(c *fiber.Ctx) (authdto.ReadTokenRes, error) {

	var (
		ctx      = c.Context()
		log      = c.Locals(contextkey.Logger).(*zap.Logger)
		claim, _ = c.Locals(contextkey.User).(*jwt.AuthClaims)
		data     authdto.ReadTokenRes
	)

	user, err := s.authrepository.FetchUser(ctx, authquery.FetchUser{
		ID: claim.UserID,
	})
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Error(err.Error())
			return data, model.ErrInternalServer("Read token failed")
		}
		return data, model.ErrBadRequest("User not found")
	}

	data = authdto.ReadTokenRes{
		ID:   user.ID,
		Name: user.Name,
		Role: user.Role,
	}

	return data, nil
}
