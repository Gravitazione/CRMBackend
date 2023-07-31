package models

import (
	"time"
)

type ArticleTierMapping struct {
	id         uint `gorm:"primary_key"`
	article_id uint
	tier_id    uint
	created_dt time.Time
	updated_dt time.Time
	is_deleted bool
}
