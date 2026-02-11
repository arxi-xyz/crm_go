package userRepository

import (
	"crm_go/entities"
	"database/sql"
)

type Query interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

type UserRepository struct {
	db Query
}

func New(query Query) *UserRepository {
	return &UserRepository{db: query}
}

func (r *UserRepository) GetUserByPhone(phone string) (*entities.User, error) {
	row := r.db.QueryRow(
		`SELECT id, uuid, phone, first_name, last_name, password,created_at, updated_at
		 FROM users
		 WHERE phone = $1 AND deleted_at IS NULL`,
		phone,
	)

	var user entities.User
	err := row.Scan(
		&user.ID,
		&user.UUID,
		&user.Phone,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
