package userService

import "crm_go/entities"

type UserService struct {
	UserRepository userRepositoryInterface
}

func New(userRepository userRepositoryInterface) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

type userRepositoryInterface interface {
	GetUserBy(field, value string) (*entities.User, error)
}
