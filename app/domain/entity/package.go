package entity

import (
	"time"
)

type Package struct {
	ID          int64      `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Act         string     `json:"act" db:"act"`
	Price       int64      `json:"price" db:"price"`
	ValidMonths int        `json:"valid_months" db:"valid_months"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}
