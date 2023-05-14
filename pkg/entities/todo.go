package entities

import (
	"time"
)

type Todo struct {
	ID                uint `gorm:"primarykey;column:todo_id" json:"id"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Activity_group_id uint   `json:"activity_group_id"`
	Title             string `json:"title"`
	Is_active         bool   `json:"is_active"`
	Priority          string `json:"priority"`
}
