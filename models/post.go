package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int    `json:"author_id"`
	Tags     []Tag  `json:"tags"`
}
