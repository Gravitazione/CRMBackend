package models

import (
	"time"
)

type Article struct {
	id              uint `gorm:"primary_key"`
	reference       string
	title           string
	description     string
	content_th      string
	content_en      string
	seq             uint
	cover_image_url string
	status          string
	valid_from      time.Time
	valid_through   time.Time
	created_dt      time.Time
	updated_dt      time.Time
	is_deleted      bool
}
