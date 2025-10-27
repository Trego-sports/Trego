package models

import (
	"time"
)

// Game represents a game/event in the system
type Game struct {
	GameID      string         `json:"game_id" db:"game_id"`
	HostID      string         `json:"host_id" db:"host_id"`
	SportName   string         `json:"sport_name" db:"sport_name"`
	Title       string         `json:"title" db:"title"`
	Description *string        `json:"description,omitempty" db:"description"`
	StartTime   time.Time      `json:"start_time" db:"start_time"`
	EndTime     time.Time      `json:"end_time" db:"end_time"`
	Location    string         `json:"location" db:"location"`
	SkillRange  *string        `json:"skill_range,omitempty" db:"skill_range"`
	Capacity    int            `json:"capacity" db:"capacity"`
	SkillLevel  *string        `json:"skill_level,omitempty" db:"skill_level"` // "beginner", "intermediate", "advanced"
	Visibility  string         `json:"visibility" db:"visibility"`             // "public" or "invite-only"
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
	Host        *User          `json:"host,omitempty"`
	Sport       *Sport         `json:"sport,omitempty"`
	Players     []GamePlayer   `json:"players,omitempty"`
	PlayerCount int            `json:"player_count,omitempty"`
}

// GamePlayer represents the many-to-many relationship between games and users (players)
type GamePlayer struct {
	UserID     string    `json:"user_id" db:"user_id"`
	GameID     string    `json:"game_id" db:"game_id"`
	Attendance string    `json:"attendance" db:"attendance"` // "true", "false", "none"
	JoinedAt   time.Time `json:"joined_at" db:"joined_at"`
	User       *User     `json:"user,omitempty"`
}

// CreateGameRequest represents the request payload for creating a new game
type CreateGameRequest struct {
	SportName   string     `json:"sport_name" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	Description *string    `json:"description,omitempty"`
	StartTime   time.Time  `json:"start_time" binding:"required"`
	EndTime     time.Time  `json:"end_time" binding:"required"`
	Location    string     `json:"location" binding:"required"`
	SkillRange  *string    `json:"skill_range,omitempty"`
	Capacity    int        `json:"capacity" binding:"required,min=1"`
	SkillLevel  *string    `json:"skill_level,omitempty" binding:"omitempty,oneof=beginner intermediate advanced"`
	Visibility  string     `json:"visibility" binding:"required,oneof=public invite-only"`
}

// UpdateGameRequest represents the request payload for updating a game
type UpdateGameRequest struct {
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	StartTime   *time.Time `json:"start_time,omitempty"`
	EndTime     *time.Time `json:"end_time,omitempty"`
	Location    *string    `json:"location,omitempty"`
	SkillRange  *string    `json:"skill_range,omitempty"`
	Capacity    *int       `json:"capacity,omitempty" binding:"omitempty,min=1"`
	SkillLevel  *string    `json:"skill_level,omitempty" binding:"omitempty,oneof=beginner intermediate advanced"`
	Visibility  *string    `json:"visibility,omitempty" binding:"omitempty,oneof=public invite-only"`
}

// JoinGameRequest represents the request payload for joining a game
type JoinGameRequest struct {
	Attendance string `json:"attendance" binding:"required,oneof=true false none"`
}

// GameFilters represents filters for querying games
type GameFilters struct {
	SportName   *string    `json:"sport_name,omitempty"`
	Location    *string    `json:"location,omitempty"`
	SkillLevel  *string    `json:"skill_level,omitempty"`
	Visibility  *string    `json:"visibility,omitempty"`
	StartAfter  *time.Time `json:"start_after,omitempty"`
	StartBefore *time.Time `json:"start_before,omitempty"`
	HostID      *string    `json:"host_id,omitempty"`
	Limit       int        `json:"limit,omitempty"`
	Offset      int        `json:"offset,omitempty"`
}
