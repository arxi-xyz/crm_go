package userRepository

import (
	"crm_go/entities"
	"database/sql"
	"fmt"
)

type Query interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

type UserRepository struct {
	db    Query
	table string
}

func New(query Query) *UserRepository {
	return &UserRepository{db: query, table: "users"}
}

func (r *UserRepository) GetUserBy(field, value string) (*entities.User, error) {
	query := fmt.Sprintf(`
	SELECT id, uuid, phone, first_name, last_name, password, created_at, updated_at
	FROM %s
	WHERE %s = $1 AND deleted_at IS NULL
`, r.table, field)

	row := r.db.QueryRow(query, value)

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
