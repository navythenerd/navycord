package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	db       *gorm.DB
	commands map[string]*ApplicationCommand
	config   *Config
}

func New(cfg *Config) (*Service, error) {
	// connect to db
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Berlin", cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// migrate db types
	err = db.AutoMigrate(&Message{})

	if err != nil {
		return nil, err
	}

	s := &Service{
		db:       db,
		commands: map[string]*ApplicationCommand{},
		config:   cfg,
	}

	return s, nil
}

func (s *Service) DB() *gorm.DB {
	return s.db
}
