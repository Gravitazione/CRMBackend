package models

import (
	"time"
)

type Customer struct {
	id         uint `gorm:"primary_key"`
	name       string
	email      string
	password   string
	created_dt time.Time
	updated_dt time.Time
	is_deleted bool
	role_id    uint
}
