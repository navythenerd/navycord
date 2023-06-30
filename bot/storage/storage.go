package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db     *gorm.DB
	config *Config
}

func New(cfg *Config) (*Storage, error) {
	// connect to db
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Berlin", cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// migrate db types
	err = db.AutoMigrate()

	if err != nil {
		return nil, err
	}

	storage := &Storage{
		db:     db,
		config: cfg,
	}

	return storage, nil
}

func (s *Storage) DB() *gorm.DB {
	return s.db
}
