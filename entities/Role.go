package entities

import "time"

type Role struct {
	ID        int
	UUID      string
	Title     string
	ParentID  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
