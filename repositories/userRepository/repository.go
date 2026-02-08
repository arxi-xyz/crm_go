package userRepository

import (
	"crm_go/entities"
	"crm_go/pkg/appError"
)

type UserRepository struct {
}

func (r *UserRepository) GetUserByPhone(phone string) (entities.User, error) {
	// todo: refactor
	if phone == "09130109810" {
		return entities.User{}, nil
	}
	return entities.User{}, appError.ErrUserNotFound
}
