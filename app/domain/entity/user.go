package entity

import (
	"time"
)

type User struct {
	ID         int64      `json:"id" db:"id"`
	Name       string     `json:"name" db:"name"`
	Email      string     `json:"email" db:"email"`
	Password   string     `json:"password" db:"password"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at"`
	Birthdate  *time.Time `json:"birthdate" db:"birthdate"`
	Location   *string    `json:"location" db:"location"`
	Image      *string    `json:"image" db:"image"`
	Gender     int        `json:"gender" db:"gender"`
	GenderLang string     `json:"gender_lang" db:"gender_lang"`
	Label      *string    `json:"label" db:"label"`
}
