package models

type Role struct {
	id    uint `gorm:"primary_key"`
	title string
}
