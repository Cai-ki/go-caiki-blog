package models

import (
	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	Content string
	PostID  uint
	UserID  uint
	User    Users `gorm:"foreignKey:UserID"`
	Post    Posts `gorm:"foreignKey:PostID"`
}
