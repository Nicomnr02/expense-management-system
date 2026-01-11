package mocks

import (
	"context"

	authdomain "expense-management-system/cmd/auth/domain"
	authquery "expense-management-system/cmd/auth/repository/query"

	"github.com/stretchr/testify/mock"
)

type AuthRepositoryMock struct {
	mock.Mock
}

func (m *AuthRepositoryMock) FetchUser(
	c context.Context,
	q authquery.FetchUser,
) (authdomain.User, error) {

	args := m.Called(c, q)

	return args.Get(0).(authdomain.User), args.Error(1)
}
