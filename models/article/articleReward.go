package models

import (
	"time"
)

type ArticleReward struct {
	id           uint `gorm:"primary_key"`
	article_id   uint
	type_id      uint
	condition_id uint
	created_dt   time.Time
	updated_dt   time.Time
	is_deleted   bool
}
