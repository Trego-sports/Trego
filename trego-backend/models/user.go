package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	UserID       string     `json:"user_id" db:"user_id"`
	Name         string     `json:"name" db:"name"`
	Email        string     `json:"email" db:"email"`
	PictureURL   *string    `json:"picture_url,omitempty" db:"picture_url"`
	PhoneNumber  *string    `json:"phone_number,omitempty" db:"phone_number"`
	Location     *string    `json:"location,omitempty" db:"location"`
	Reputation   int        `json:"reputation" db:"reputation"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	Sports       []UserSport `json:"sports,omitempty"`
}

// UserSport represents the many-to-many relationship between users and sports
type UserSport struct {
	UserID      string    `json:"user_id" db:"user_id"`
	SportName   string    `json:"sport_name" db:"sport_name"`
	Position    *string   `json:"position,omitempty" db:"position"` // "front" or "back"
	SkillLevel  string    `json:"skill_level" db:"skill_level"`     // "beginner", "intermediate", "advanced"
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Sport       *Sport    `json:"sport,omitempty"`
}

// CreateUserRequest represents the request payload for creating a new user
type CreateUserRequest struct {
	Name        string  `json:"name" binding:"required"`
	Email       string  `json:"email" binding:"required,email"`
	PictureURL  *string `json:"picture_url,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Location    *string `json:"location,omitempty"`
}

// UpdateUserRequest represents the request payload for updating a user
type UpdateUserRequest struct {
	Name        *string `json:"name,omitempty"`
	PictureURL  *string `json:"picture_url,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Location    *string `json:"location,omitempty"`
}

// AddUserSportRequest represents the request payload for adding a sport to a user
type AddUserSportRequest struct {
	SportName   string  `json:"sport_name" binding:"required"`
	Position    *string `json:"position,omitempty"`
	SkillLevel  string  `json:"skill_level" binding:"required,oneof=beginner intermediate advanced"`
}

// UpdateUserSportRequest represents the request payload for updating a user's sport
type UpdateUserSportRequest struct {
	Position   *string `json:"position,omitempty"`
	SkillLevel *string `json:"skill_level,omitempty"`
}
