package models

import (
	"time"
)

type Pose struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	PoseName        string         `gorm:"type:varchar(255);not null" json:"pose_name"`
	PoseImage       []byte         `gorm:"type:bytea" json:"pose_image,omitempty"` // Storing Blob
	PoseDescription string         `gorm:"type:text" json:"pose_description"`
	PoseCondition 	string         `gorm:"type:text" json:"pose_condition"`
	PoseCategoryID  uint           `gorm:"not null" json:"pose_category_id"`
	PoseAccuracy    float64        `gorm:"type:float" json:"pose_accuracy"`
	Status          string         `gorm:"type:varchar(50)" json:"status"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	
	PoseCategory    PoseCategory   `gorm:"foreignKey:PoseCategoryID" json:"pose_category"`
}
