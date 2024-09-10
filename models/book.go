package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID            uint           `json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Title         string         `json:"title"`
	AuthorId      uint           `json:"author_id"`
	PublishedYear int            `json:"published_year"`
}
