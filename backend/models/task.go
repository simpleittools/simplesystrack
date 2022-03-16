package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	TaskName    string         `json:"task_name"`
	TaskCode    string         `json:"task_code" gorm:"unique"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	IsActive    *bool          `json:"is_active" gorm:"default:true"`
	TaskDetails string         `json:"task_details" gorm:"type:text" gorm:"nullable"`
	MilestoneID int            `json:"milestone_id"`
	Milestone   Milestone      `json:"milestone"`
}

// Linking for tasks and milestones is not working
