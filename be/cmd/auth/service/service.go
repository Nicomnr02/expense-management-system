package authservice

import authrepository "expense-management-system/cmd/auth/repository"

type AuthService interface{
	
}

type authserviceImpl struct {
	authrepository authrepository.Authrepository
}

func New(authrepository authrepository.Authrepository) AuthService {
	return authserviceImpl{authrepository}
}

