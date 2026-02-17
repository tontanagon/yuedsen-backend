package models

import (
	"time"
)

type UserProcess struct {
	ID             uint         `gorm:"primaryKey" json:"id"`
	UserID         uint         `gorm:"not null" json:"user_id"`
	PoseCategoryID uint         `gorm:"not null" json:"pose_category_id"`
	Progress       int          `gorm:"type:int;default:1" json:"progress"` // Progress percentage
	Status         string       `gorm:"type:varchar(50)" json:"status"`   // e.g., "in_progress", "completed"
	TotalScore     int          `gorm:"type:int" json:"total_score"`      // Cumulative score
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`

	User         User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
	PoseCategory PoseCategory `gorm:"foreignKey:PoseCategoryID" json:"pose_category,omitempty"`
}
