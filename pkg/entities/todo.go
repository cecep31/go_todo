package entities

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID                uuid.UUID `gorm:"primarykey;column:todo_id" json:"id"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Activity_group_id uint   `json:"activity_group_id"`
	Title             string `json:"title"`
	Is_active         bool   `json:"is_active"`
	Priority          string `json:"priority"`
}
