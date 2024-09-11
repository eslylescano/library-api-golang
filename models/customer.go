package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
}
