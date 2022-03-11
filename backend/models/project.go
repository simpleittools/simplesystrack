package models

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ProjectName string         `json:"project_name"`
	ProjectCode string         `json:"project_code" gorm:"unique"`
	ClientID    int            `json:"client_id"`
	Client      Client         `json:"client"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Milestones  []Milestone
}
