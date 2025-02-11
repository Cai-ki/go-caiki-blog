package models

import (
	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	PostID  int    `json:"post_id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}
