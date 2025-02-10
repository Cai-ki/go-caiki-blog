package storage

import "gorm.io/gorm"

type Storage interface {
}

type storageImpl struct {
	db *gorm.DB
}

var _ Storage = (*storageImpl)(nil)

var DB Storage

func SetupStorage(db *gorm.DB) {
	DB = &storageImpl{db: db}
}
