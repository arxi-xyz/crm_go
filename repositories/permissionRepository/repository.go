package permissionRepository

import (
	"crm_go/entities"
	"crm_go/pkg/appError"
	"database/sql"
)

var fields = []string{
	"id", "uuid", "unique_key", "title",
}

type Query interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...any) (*sql.Rows, error)
}

type PermissionRepository struct {
	db    Query
	table string
}

func New(query Query) *PermissionRepository {
	return &PermissionRepository{db: query, table: "permissions"}
}

func (r *PermissionRepository) GetUserPermissions(userId int) ([]entities.Permission, *appError.AppError) {
	query := `
	SELECT p.*
	from model_has_permission AS mhp
	JOIN permissions AS p
		ON p.id = mhp.permission_id
	where mhp.model_id = $1
	AND mhp.model_type = 'users'
	`

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, appError.Internal(err)
	}
	defer rows.Close()

	var permissions []entities.Permission

	for rows.Next() {
		var p entities.Permission

		err := rows.Scan(
			&p.ID,
			&p.UUID,
			&p.Title,
			&p.UniqueKey,
		)
		if err != nil {
			return nil, appError.Internal(err)
		}

		permissions = append(permissions, p)
	}

	if err := rows.Err(); err != nil {
		return nil, appError.Internal(err)
	}

	return permissions, nil
}
