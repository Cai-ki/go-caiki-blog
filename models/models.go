package models

import "gorm.io/gorm"

func SetupModels(db *gorm.DB) {
	db.AutoMigrate(&Users{})
	// db.AutoMigrate(&Posts{})
	// db.AutoMigrate(&Tags{})
	// db.AutoMigrate(&Posts{})
}
