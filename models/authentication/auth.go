package models

import (
	"time"
)

type Auth struct {
	id            uint `gorm:"primary_key"`
	user_id       uint
	access_token  string
	refresh_token string
	created_dt    time.Time
	updated_dt    time.Time
	is_deleted    bool
}
