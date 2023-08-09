package storage

import "time"

type Message struct {
	ID         string `gorm:"primaryKey"`
	User       string
	Content    string
	Annotation string
	Hash       string
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
