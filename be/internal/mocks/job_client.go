package mocks

import (
	"expense-management-system/internal/job"

	"github.com/stretchr/testify/mock"
)

type JobClientMock struct {
	mock.Mock
}

func (m *JobClientMock) Enqueue(t job.Task) error {
	args := m.Called(t)
	return args.Error(0)
}

func (m *JobClientMock) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *JobClientMock) Ping() error {
	args := m.Called()
	return args.Error(0)
}
