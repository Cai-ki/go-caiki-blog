package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null;unique"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"` // 密码需加密存储
	Posts    []Posts
	// AvatarURL string
}
