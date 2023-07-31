package models

import (
	"time"
)

type ArticleCategory struct {
	id                 uint `gorm:"primary_key"`
	reference          string
	name_th            string
	name_en            string
	seq                uint
	image_active_url   string
	image_inactive_url string
	status             string
	valid_from         time.Time
	valid_through      time.Time
	created_dt         time.Time
	updated_dt         time.Time
	is_deleted         bool
}
