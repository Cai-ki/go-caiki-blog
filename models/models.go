package models

import "gorm.io/gorm"

func SetupModels(db *gorm.DB) {
	db.AutoMigrate(&User{})
	// db.AutoMigrate(&Post{})
	// db.AutoMigrate(&Tag{})
	// db.AutoMigrate(&Post{})
}
