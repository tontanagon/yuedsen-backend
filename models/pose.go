package models

import (
	"time"
)

type Pose struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	PoseName        string         `gorm:"type:varchar(255);not null" json:"pose_name"`
	PoseImage       string         `gorm:"type:text" json:"pose_image,omitempty"` // Storing URL or Base64 Text
	PoseDescription string         `gorm:"type:text" json:"pose_description"`
	PoseCondition 	string         `gorm:"type:text" json:"pose_condition"`

	// New fields for specific points
	PosePoint       string          `gorm:"type:text" json:"pose_point"` // JSON string: [11, 13, 15] 
	
	Status          string         `gorm:"type:varchar(50)" json:"status"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`

	// Has-many relationship for landmark-based comparison
	Landmarks []PoseLandmark `gorm:"foreignKey:PoseID" json:"landmarks,omitempty"`
}
