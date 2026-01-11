package mocks

import "github.com/stretchr/testify/mock"

type ValidatorMock struct {
	mock.Mock
}

func (m *ValidatorMock) ValidateStruct(s interface{}) error {
	args := m.Called(s)
	return args.Error(0)
}
