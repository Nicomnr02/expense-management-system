package mocks

import "expense-management-system/config"

type Mocks struct {
	Authrepo    *AuthRepositoryMock
	Expenserepo *ExpenseRepositoryMock
	Transaction *TransactionMock
	Validator   *ValidatorMock
	JobClient   *JobClientMock
	Config      *config.Config
}

func New() Mocks {
	authrepo := new(AuthRepositoryMock)
	expenserepo := new(ExpenseRepositoryMock)
	transaction := new(TransactionMock)
	validator := new(ValidatorMock)
	jobClient := new(JobClientMock)
	config := Config()

	return Mocks{
		Authrepo:    authrepo,
		Expenserepo: expenserepo,
		Transaction: transaction,
		Validator:   validator,
		JobClient:   jobClient,
		Config:      config,
	}
}
