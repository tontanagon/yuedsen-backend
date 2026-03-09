package models

type PoseLandmark struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	PoseID        uint    `gorm:"not null" json:"pose_id"`
	LandmarkIndex int     `gorm:"not null" json:"landmark_index"` // 0-32
	X             float64 `gorm:"type:float;not null" json:"x"`
	Y             float64 `gorm:"type:float;not null" json:"y"`
	Z             float64 `gorm:"type:float;not null" json:"z"`
	Visibility    float64 `gorm:"type:float" json:"visibility"`
	
	Pose Pose `gorm:"foreignKey:PoseID" json:"-"`
}
