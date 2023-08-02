package models

import (
	"time"
)

type Tier struct {
	id          uint `gorm:"primary_key"`
	name        string
	description string
	created_dt  time.Time
	updated_dt  time.Time
	is_deleted  bool
}
