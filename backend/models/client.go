package models

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	ID         uint   `json:"id"`
	ClientName string `json:"client_name"`
	ClientID   string `json:"client_id"`
	//Email     string         `json:"email" gorm:"unique"`
	//Password  []byte         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
