package models

import "gorm.io/gorm"

type Tags struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(255);not null;unique;index"`
	Posts []Posts `gorm:"many2many:posts_tags;"`
}
