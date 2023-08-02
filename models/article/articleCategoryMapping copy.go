package models

type ArticleCategoryMapping struct {
	id          uint `gorm:"primary_key"`
	category_id uint
	article_id  uint
	is_deleted  bool
}
