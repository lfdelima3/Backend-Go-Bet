package model

import (
	"time"
)

type MatchEvent struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	MatchID     uint   `json:"match_id" validate:"required"`
	EventType   string `json:"event_type" validate:"required,oneof=goal card foul substitution throw_in corner"`
	TeamID      uint   `json:"team_id" validate:"required"`
	PlayerID    uint   `json:"player_id" validate:"required"`
	Minute      int    `json:"minute" validate:"required,min=1,max=120"`
	Description string `json:"description" validate:"required,min=3,max=500"`
	// Campos específicos para cada tipo de evento
	GoalType       string    `json:"goal_type,omitempty" validate:"omitempty,oneof=normal penalty own_goal"`
	CardType       string    `json:"card_type,omitempty" validate:"omitempty,oneof=yellow red"`
	FoulType       string    `json:"foul_type,omitempty" validate:"omitempty,oneof=normal dangerous serious"`
	SubInPlayerID  uint      `json:"sub_in_player_id,omitempty" validate:"omitempty"`
	SubOutPlayerID uint      `json:"sub_out_player_id,omitempty" validate:"omitempty"`
	ThrowInSide    string    `json:"throw_in_side,omitempty" validate:"omitempty,oneof=left right"`
	CornerSide     string    `json:"corner_side,omitempty" validate:"omitempty,oneof=left right"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type MatchEventCreate struct {
	MatchID        uint   `json:"match_id" validate:"required"`
	EventType      string `json:"event_type" validate:"required,oneof=goal card foul substitution throw_in corner"`
	TeamID         uint   `json:"team_id" validate:"required"`
	PlayerID       uint   `json:"player_id" validate:"required"`
	Minute         int    `json:"minute" validate:"required,min=1,max=120"`
	Description    string `json:"description" validate:"required,min=3,max=500"`
	GoalType       string `json:"goal_type,omitempty" validate:"omitempty,oneof=normal penalty own_goal"`
	CardType       string `json:"card_type,omitempty" validate:"omitempty,oneof=yellow red"`
	FoulType       string `json:"foul_type,omitempty" validate:"omitempty,oneof=normal dangerous serious"`
	SubInPlayerID  uint   `json:"sub_in_player_id,omitempty" validate:"omitempty"`
	SubOutPlayerID uint   `json:"sub_out_player_id,omitempty" validate:"omitempty"`
	ThrowInSide    string `json:"throw_in_side,omitempty" validate:"omitempty,oneof=left right"`
	CornerSide     string `json:"corner_side,omitempty" validate:"omitempty,oneof=left right"`
}

type MatchEventUpdate struct {
	EventType      string `json:"event_type,omitempty" validate:"omitempty,oneof=goal card foul substitution throw_in corner"`
	TeamID         uint   `json:"team_id,omitempty" validate:"omitempty"`
	PlayerID       uint   `json:"player_id,omitempty" validate:"omitempty"`
	Minute         int    `json:"minute,omitempty" validate:"omitempty,min=1,max=120"`
	Description    string `json:"description,omitempty" validate:"omitempty,min=3,max=500"`
	GoalType       string `json:"goal_type,omitempty" validate:"omitempty,oneof=normal penalty own_goal"`
	CardType       string `json:"card_type,omitempty" validate:"omitempty,oneof=yellow red"`
	FoulType       string `json:"foul_type,omitempty" validate:"omitempty,oneof=normal dangerous serious"`
	SubInPlayerID  uint   `json:"sub_in_player_id,omitempty" validate:"omitempty"`
	SubOutPlayerID uint   `json:"sub_out_player_id,omitempty" validate:"omitempty"`
	ThrowInSide    string `json:"throw_in_side,omitempty" validate:"omitempty,oneof=left right"`
	CornerSide     string `json:"corner_side,omitempty" validate:"omitempty,oneof=left right"`
}

type MatchEventResponse struct {
	ID             uint      `json:"id"`
	MatchID        uint      `json:"match_id"`
	EventType      string    `json:"event_type"`
	TeamID         uint      `json:"team_id"`
	PlayerID       uint      `json:"player_id"`
	Minute         int       `json:"minute"`
	Description    string    `json:"description"`
	GoalType       string    `json:"goal_type,omitempty"`
	CardType       string    `json:"card_type,omitempty"`
	FoulType       string    `json:"foul_type,omitempty"`
	SubInPlayerID  uint      `json:"sub_in_player_id,omitempty"`
	SubOutPlayerID uint      `json:"sub_out_player_id,omitempty"`
	ThrowInSide    string    `json:"throw_in_side,omitempty"`
	CornerSide     string    `json:"corner_side,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Tipos de eventos
const (
	EventTypeGoal         = "goal"
	EventTypeCard         = "card"
	EventTypeFoul         = "foul"
	EventTypeSubstitution = "substitution"
	EventTypeThrowIn      = "throw_in"
	EventTypeCorner       = "corner"
)

// Tipos de gols
const (
	GoalTypeNormal  = "normal"
	GoalTypePenalty = "penalty"
	GoalTypeOwnGoal = "own_goal"
)

// Tipos de cartões
const (
	CardTypeYellow = "yellow"
	CardTypeRed    = "red"
)

// Tipos de faltas
const (
	FoulTypeNormal    = "normal"
	FoulTypeDangerous = "dangerous"
	FoulTypeSerious   = "serious"
)
