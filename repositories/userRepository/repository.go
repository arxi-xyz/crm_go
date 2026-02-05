package userRepository

import "crm_go/entities"

type UserRepository struct {
}

func (r *UserRepository) GetUserByPhone(phone string) (entities.User, error) {
	return entities.User{}, nil
}
