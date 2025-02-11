package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255);not null;unique"`
	Email    string `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password string `json:"password" gorm:"type:varchar(255);not null"` // 密码需加密存储
	// AvatarURL string `json:"avatar_url"`
}
