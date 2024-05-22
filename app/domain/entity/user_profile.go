package entity

import (
	"database/sql"
	"time"
)

type UserProfile struct {
	ID        int64          `json:"id" db:"id"`
	UserID    int64          `json:"user_id" db:"user_id"`
	Birthdate sql.NullTime   `json:"birthdate" db:"birthdate"`
	Location  sql.NullString `json:"location" db:"location"`
	Image     sql.NullString `json:"image" db:"image"`
	Gender    int            `json:"gender" db:"gender"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime   `json:"deleted_at" db:"deleted_at"`
}
