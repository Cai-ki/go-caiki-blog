package models

import (
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint   `gorm:"not null"`
	User    Users  `gorm:"foreignKey:UserID"`
}
