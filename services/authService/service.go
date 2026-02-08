package authService

import (
	"crm_go/entities"
)

type AuthService struct {
	UserRepository userRepositoryInterface
}

func New(userRepository userRepositoryInterface) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

type userRepositoryInterface interface {
	GetUserByPhone(phone string) (entities.User, error)
}
