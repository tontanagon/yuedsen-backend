package models

import (
	"time"
)

type PoseCategory struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Category  string    `gorm:"type:varchar(255);unique;not null" json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
