package storage

import (
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New() *Storage {
	// Инициализация подключения к БД
	return &Storage{}
}

func (s *Storage) Create(data interface{}) error {
	// Реализация сохранения в БД
	return nil
}