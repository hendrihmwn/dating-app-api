package entity

import (
	"time"
)

type UserPair struct {
	ID         int64      `json:"id" db:"id"`
	UserID     int64      `json:"user_id" db:"user_id"`
	PairUserID int64      `json:"pair_user_id" db:"pair_user_id"`
	Status     int        `json:"status" db:"status"`
	StatusLang string     `json:"status_lang" db:"status_lang"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at"`
}
