package models

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	ClientName     string         `json:"client_name"`
	ClientInitials string         `json:"client_initials"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	Projects       []Project
}
