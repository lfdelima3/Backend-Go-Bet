package model

import (
	"time"
)

type Corner struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	MatchID     uint      `json:"match_id" validate:"required"`
	TeamID      uint      `json:"team_id" validate:"required"`
	PlayerID    uint      `json:"player_id" validate:"required"`
	Minute      int       `json:"minute" validate:"required,min=1,max=120"`
	Side        string    `json:"side" validate:"required,oneof=left right"`
	Result      string    `json:"result" validate:"required,oneof=goal chance cleared saved"`
	Description string    `json:"description" validate:"omitempty,max=200"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CornerCreate struct {
	MatchID     uint   `json:"match_id" validate:"required"`
	TeamID      uint   `json:"team_id" validate:"required"`
	PlayerID    uint   `json:"player_id" validate:"required"`
	Minute      int    `json:"minute" validate:"required,min=1,max=120"`
	Side        string `json:"side" validate:"required,oneof=left right"`
	Result      string `json:"result" validate:"required,oneof=goal chance cleared saved"`
	Description string `json:"description" validate:"omitempty,max=200"`
}

type CornerUpdate struct {
	TeamID      uint   `json:"team_id" validate:"omitempty"`
	PlayerID    uint   `json:"player_id" validate:"omitempty"`
	Minute      int    `json:"minute" validate:"omitempty,min=1,max=120"`
	Side        string `json:"side" validate:"omitempty,oneof=left right"`
	Result      string `json:"result" validate:"omitempty,oneof=goal chance cleared saved"`
	Description string `json:"description" validate:"omitempty,max=200"`
}

type CornerResponse struct {
	ID          uint      `json:"id"`
	MatchID     uint      `json:"match_id"`
	TeamID      uint      `json:"team_id"`
	PlayerID    uint      `json:"player_id"`
	Minute      int       `json:"minute"`
	Side        string    `json:"side"`
	Result      string    `json:"result"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
