package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"` // 密码需加密存储
	// AvatarURL string `json:"avatar_url"`
}
