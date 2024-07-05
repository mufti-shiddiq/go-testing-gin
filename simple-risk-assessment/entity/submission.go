package entity

import "time"

type Submission struct {
	ID        int       `gorm:"primary_key" json:"id"`
	UserID    int      `gorm:"index" json:"user_id"`
	Answers   string    `json:"answers"`
	RiskScore int       `json:"risk_score"`
	RiskCategory string `json:"risk_category"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Submission) TableName() string {
	return "submissions"
}
