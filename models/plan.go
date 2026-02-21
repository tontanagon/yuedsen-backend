package models

import (
	"time"
)

type Plan struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Day       int       `gorm:"not null" json:"day"` // Create a plan for days 1-30
	PoseID    uint      `gorm:"not null" json:"pose_id"`
	Duration     int       `gorm:"not null" json:"duration"` // Duration in seconds or number of repetitions
	PoseAccuracy float64   `gorm:"type:float" json:"pose_accuracy"`
	PoseMin      float64   `gorm:"type:float" json:"pose_min"`
	PoseMax      float64   `gorm:"type:float" json:"pose_max"`
	PoseCategoryID uint    `gorm:"not null" json:"pose_category_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Pose Pose `gorm:"foreignKey:PoseID" json:"pose"`
	PoseCategory PoseCategory `gorm:"foreignKey:PoseCategoryID" json:"pose_category"`
}
