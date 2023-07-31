package models

import (
	"time"
)

type ArticleRewardCriteria struct {
	id            uint `gorm:"primary_key"`
	article_id    uint
	tier_id       uint
	inventory_ref string
	privilege_ref string
	point         uint
	created_dt    time.Time
	updated_dt    time.Time
	is_deleted    bool
}
