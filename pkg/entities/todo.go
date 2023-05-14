package entities

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Activity_group_id uint   `json:"activity_group_id"`
	Title             string `json:"title"`
	Is_active         bool   `json:"is_active"`
	Priority          string `json:"priority"`
}
