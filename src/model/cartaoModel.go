package model

import (
	"time"
)

type Card struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	MatchID     uint      `json:"match_id" validate:"required"`
	TeamID      uint      `json:"team_id" validate:"required"`
	PlayerID    uint      `json:"player_id" validate:"required"`
	Minute      int       `json:"minute" validate:"required,min=1,max=120"`
	CardType    string    `json:"card_type" validate:"required,oneof=yellow red"`
	Reason      string    `json:"reason" validate:"required,max=200"`
	Description string    `json:"description" validate:"omitempty,max=200"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CardCreate struct {
	MatchID     uint   `json:"match_id" validate:"required"`
	TeamID      uint   `json:"team_id" validate:"required"`
	PlayerID    uint   `json:"player_id" validate:"required"`
	Minute      int    `json:"minute" validate:"required,min=1,max=120"`
	CardType    string `json:"card_type" validate:"required,oneof=yellow red"`
	Reason      string `json:"reason" validate:"required,max=200"`
	Description string `json:"description" validate:"omitempty,max=200"`
}

type CardUpdate struct {
	TeamID      uint   `json:"team_id" validate:"omitempty"`
	PlayerID    uint   `json:"player_id" validate:"omitempty"`
	Minute      int    `json:"minute" validate:"omitempty,min=1,max=120"`
	CardType    string `json:"card_type" validate:"omitempty,oneof=yellow red"`
	Reason      string `json:"reason" validate:"omitempty,max=200"`
	Description string `json:"description" validate:"omitempty,max=200"`
}

type CardResponse struct {
	ID          uint      `json:"id"`
	MatchID     uint      `json:"match_id"`
	TeamID      uint      `json:"team_id"`
	PlayerID    uint      `json:"player_id"`
	Minute      int       `json:"minute"`
	CardType    string    `json:"card_type"`
	Reason      string    `json:"reason"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
