package entities

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	UUID      string
	Phone     string
	Password  string
	FirstName sql.NullString
	LastName  sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}
