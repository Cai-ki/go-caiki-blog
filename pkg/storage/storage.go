package storage

import "gorm.io/gorm"

type Storage interface {
	GetDB() *gorm.DB
}

type storageImpl struct {
	db *gorm.DB
}

var _ Storage = (*storageImpl)(nil)

var DB Storage

func SetupStorage(db *gorm.DB) {
	DB = &storageImpl{db: db}
}

func (s *storageImpl) GetDB() *gorm.DB {
	if s.db == nil {
		panic("DB is not initialized")
	}
	return s.db
}
