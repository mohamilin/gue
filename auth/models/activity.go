package models

import (
	"time"
)

type Activity struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Service   string    `json:"service"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
	Timestamp time.Time `json:"timestamp" sql:"DEFAULT:current_timestamp"`
}
