package models

import (
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int    `json:"author_id"`
	Tags     []Tags `json:"tags"`
}
