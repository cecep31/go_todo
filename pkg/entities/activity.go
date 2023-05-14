package entities

import (
	"time"
)

type Activity struct {
	ID        uint `gorm:"primarykey;column:activity_id" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string `json:"title"`
	Email     string `json:"email"`
}
