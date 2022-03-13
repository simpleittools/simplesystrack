package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Milestone struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	MilestoneName     string         `json:"milestone_name"`
	MilestoneCode     string         `json:"milestone_code" gorm:"unique"`
	MilestoneStarted  *bool          `json:"milestone_started" gorm:"default:false"`
	MilestonePrereq   int            `json:"milestone_prereq" gorm:"nullable"`
	MilestoneStart    datatypes.Date `json:"milestone_start" gorm:"nullable"`
	MilestoneEnd      datatypes.Date `json:"milestone_end" gorm:"nullable"`
	MilestoneDetails  string         `json:"milestone_details"`
	MilestoneIsActive *bool          `json:"milestone_is_active" gorm:"default:false"`
	ProjectID         int            `json:"project_id"`
	Project           Project        `json:"project"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at"`
}
