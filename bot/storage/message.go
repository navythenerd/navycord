package storage

import "time"

type Message struct {
	ID         uint
	User       string
	Content    string
	Annotation string
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
