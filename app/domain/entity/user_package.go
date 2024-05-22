package entity

import (
	"time"
)

type UserPackage struct {
	ID         int64      `json:"id" db:"id"`
	UserID     int64      `json:"user_id" db:"user_id"`
	PackageID  int64      `json:"package_id" db:"package_id"`
	PackageAct string     `json:"package_act" db:"package_act"`
	ExpiredAt  *time.Time `json:"expired_at" db:"expired_at"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at"`
}
