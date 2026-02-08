package models

import (
	"time"
)

type Procedure struct {
	ID         int       `json:"id" db:"id"`
	Title      string    `json:"title" db:"title"`
	Type       string    `json:"type" db:"type"`
	Content    string    `json:"content" db:"content"`
	SortOrder  int       `json:"sort_order" db:"sort_order"`
	IsExpanded bool      `json:"is_expanded" db:"is_expanded"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
