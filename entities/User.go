package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int64
	UUID      uuid.UUID
	Phone     string
	FirstName sql.NullString
	LastName  sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}
