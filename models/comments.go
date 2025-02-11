package models

import (
	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	PostID  int
	UserID  int
	Content string
}
