package models

import (
	"time"
)

// Sport represents a sport in the system
type Sport struct {
	SportName  string    `json:"sport_name" db:"sport_name"`
	IconURL    *string   `json:"icon_url,omitempty" db:"icon_url"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// CreateSportRequest represents the request payload for creating a new sport
type CreateSportRequest struct {
	SportName string  `json:"sport_name" binding:"required"`
	IconURL   *string `json:"icon_url,omitempty"`
}

// UpdateSportRequest represents the request payload for updating a sport
type UpdateSportRequest struct {
	IconURL *string `json:"icon_url,omitempty"`
}
