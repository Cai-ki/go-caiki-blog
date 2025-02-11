package models

import "gorm.io/gorm"

type Tags struct {
	gorm.Model
	Name string
}
