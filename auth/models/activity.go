package models

import (
	"encoding/json"
	"time"
)

type Activity struct {
	ID        uint            `json:"id" gorm:"primaryKey"`
	UserID    uint            `json:"user_id"`
	Service   string          `json:"service"`
	Path      string          `json:"path"`
	Method    string          `json:"method"`
	Body      json.RawMessage `json:"body"`
	Query     string          `json:"query"`
	Timestamp time.Time       `json:"timestamp" sql:"DEFAULT:current_timestamp"`
}
