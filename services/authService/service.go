package authService

import (
	"crm_go/entities"
	"crm_go/repositories/userRepository"
)

type AuthService struct {
	UserRepository userRepositoryInterface
}

func New(userRepository userRepository.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: &userRepository,
	}
}

type userRepositoryInterface interface {
	GetUserByPhone(phone string) (entities.User, error)
}
