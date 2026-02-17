package models

import (
	"time"
)

type Plan struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Day       int       `gorm:"not null" json:"day"` // Create a plan for days 1-30
	PoseID    uint      `gorm:"not null" json:"pose_id"`
	Duration  int       `gorm:"not null" json:"duration"` // Duration in seconds or number of repetitions
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Pose Pose `gorm:"foreignKey:PoseID" json:"pose"`
}
