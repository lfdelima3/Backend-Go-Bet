package model

import (
	"time"
)

type Foul struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	MatchID     uint      `json:"match_id" validate:"required"`
	TeamID      uint      `json:"team_id" validate:"required"`
	PlayerID    uint      `json:"player_id" validate:"required"`
	Minute      int       `json:"minute" validate:"required,min=1,max=120"`
	FoulType    string    `json:"foul_type" validate:"required,oneof=normal dangerous serious violent"`
	Location    string    `json:"location" validate:"required,oneof=defensive midfield offensive"`
	Description string    `json:"description" validate:"omitempty,max=200"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FoulCreate struct {
	MatchID     uint   `json:"match_id" validate:"required"`
	TeamID      uint   `json:"team_id" validate:"required"`
	PlayerID    uint   `json:"player_id" validate:"required"`
	Minute      int    `json:"minute" validate:"required,min=1,max=120"`
	FoulType    string `json:"foul_type" validate:"required,oneof=normal dangerous serious violent"`
	Location    string `json:"location" validate:"required,oneof=defensive midfield offensive"`
	Description string `json:"description" validate:"omitempty,max=200"`
}

type FoulUpdate struct {
	TeamID      uint   `json:"team_id" validate:"omitempty"`
	PlayerID    uint   `json:"player_id" validate:"omitempty"`
	Minute      int    `json:"minute" validate:"omitempty,min=1,max=120"`
	FoulType    string `json:"foul_type" validate:"omitempty,oneof=normal dangerous serious violent"`
	Location    string `json:"location" validate:"omitempty,oneof=defensive midfield offensive"`
	Description string `json:"description" validate:"omitempty,max=200"`
}

type FoulResponse struct {
	ID          uint      `json:"id"`
	MatchID     uint      `json:"match_id"`
	TeamID      uint      `json:"team_id"`
	PlayerID    uint      `json:"player_id"`
	Minute      int       `json:"minute"`
	FoulType    string    `json:"foul_type"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
