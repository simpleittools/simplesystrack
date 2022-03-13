package models

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	ID                 uint           `json:"id" gorm:"primaryKey"`
	ProjectName        string         `json:"project_name"`
	ProjectCode        string         `json:"project_code" gorm:"unique"`
	ProjectDescription string         `json:"project_description" gorm:"type:text; nullable"`
	ProjectBudgetHrs   int            `json:"project_budget_hrs" gorm:"nullable"`
	ProjectBudget      int            `json:"project_budget" gorm:"nullable"`
	ProjectIsActive    *bool          `json:"project_is_active" gorm:"default:false"`
	ProjectComments    string         `json:"project_comments" gorm:"type:text; nullable"`
	ClientID           int            `json:"client_id"`
	Client             Client         `json:"client"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at"`
	Milestones         []Milestone
}
