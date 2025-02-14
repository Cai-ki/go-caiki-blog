package models

import (
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null;index"`
	Content string `gorm:"type:text;not null"`
	UserID  uint   `gorm:"not null"`
	User    Users  `gorm:"foreignKey:UserID"`
	Tags    []Tags `gorm:"many2many:posts_tags;"`
}
