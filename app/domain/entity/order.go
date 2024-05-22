package entity

import (
	"database/sql"
	"time"
)

type Order struct {
	ID           int64        `json:"id" db:"id"`
	UserID       int64        `json:"user_id" db:"user_id"`
	PackageName  string       `json:"package_name" db:"package_name"`
	PackagePrice int64        `json:"package_price" db:"package_price"`
	Status       int          `json:"status" db:"status"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
