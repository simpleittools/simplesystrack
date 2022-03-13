package models

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	ClientName     string         `json:"client_name"`
	ClientInitials string         `json:"client_initials" gorm:"unique"`
	ClientStreet   string         `json:"client_street" gorm:"nullable"`
	ClientStreet2  string         `json:"client_street_2" gorm:"nullable"`
	ClientZip      string         `json:"client_zip" gorm:"nullable"`
	ClientState    string         `json:"client_state" gorm:"nullable"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	IsActive       *bool          `json:"is_active" gorm:"default:true"`
	Comments       string         `json:"comments" gorm:"type:text"`
	Projects       []Project
}
