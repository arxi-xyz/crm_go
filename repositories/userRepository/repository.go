package userRepository

import (
	"crm_go/entities"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByPhone(phone string) (entities.User, error) {
	row := r.db.QueryRow(
		`SELECT id, uuid, phone, first_name, last_name, created_at, updated_at
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
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.User{}, nil
		}
		return entities.User{}, err
	}

	return user, nil
}
