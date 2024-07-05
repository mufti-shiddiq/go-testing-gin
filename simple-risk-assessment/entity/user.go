package entity

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Submissions []Submission `gorm:"foreignKey:UserID" json:"submissions"`
}

func (User) TableName() string {
	return "users"
}
