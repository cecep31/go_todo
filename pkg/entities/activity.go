package entities

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID        uuid.UUID `gorm:"primarykey;column:activity_id" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
}
