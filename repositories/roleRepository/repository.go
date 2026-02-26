package roleRepository

import (
	"crm_go/entities"
	"crm_go/pkg/appError"
	"database/sql"
)

var fields = []string{
	"id", "uuid", "parent_id", "title",
}

type Query interface {
	QueryRow(query string, args ...any) *sql.Row
	Query(query string, args ...any) (*sql.Rows, error)
}

type RoleRepository struct {
	db    Query
	table string
}

func New(query Query) *RoleRepository {
	return &RoleRepository{
		db:    query,
		table: "roles",
	}
}

func (r *RoleRepository) GetUserRoleWithPermissions(userId int) ([]entities.Permission, *appError.AppError) {
	query := `
	SELECT p.*
	FROM model_has_role AS mhr
	JOIN model_has_permission AS mhp
		ON mhr.role_id = mhp.model_id
		AND mhp.model_type = 'roles'
	JOIN permissions AS p
		ON p.id = mhp.permission_id
	WHERE mhr.model_id = $1
	  AND mhr.model_type = 'users'
	`

	rows, err := r.db.Query(query, userId)

	var Permissions []entities.Permission

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

		Permissions = append(Permissions, p)
	}

	if err != nil {
		return nil, appError.Internal(err)
	}

	return Permissions, nil
}
