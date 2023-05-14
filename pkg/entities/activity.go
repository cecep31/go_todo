package entities

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID        uuid.UUID `gorm:"primarykey;column:activity_id" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string `json:"title"`
	Email     string `json:"email"`
}
