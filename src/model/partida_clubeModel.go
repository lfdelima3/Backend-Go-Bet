package model

import (
	"time"
)

type MatchTeam struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	MatchID   uint      `json:"match_id" validate:"required"`
	TeamID    uint      `json:"team_id" validate:"required"`
	IsHome    bool      `json:"is_home" validate:"required"`
	Formation string    `json:"formation" validate:"required,oneof=4-4-2 4-3-3 4-2-3-1 3-5-2 5-3-2 4-5-1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MatchTeamCreate struct {
	MatchID   uint   `json:"match_id" validate:"required"`
	TeamID    uint   `json:"team_id" validate:"required"`
	IsHome    bool   `json:"is_home" validate:"required"`
	Formation string `json:"formation" validate:"required,oneof=4-4-2 4-3-3 4-2-3-1 3-5-2 5-3-2 4-5-1"`
}

type MatchTeamUpdate struct {
	TeamID    uint   `json:"team_id" validate:"omitempty"`
	IsHome    bool   `json:"is_home" validate:"omitempty"`
	Formation string `json:"formation" validate:"omitempty,oneof=4-4-2 4-3-3 4-2-3-1 3-5-2 5-3-2 4-5-1"`
}

type MatchTeamResponse struct {
	ID        uint      `json:"id"`
	MatchID   uint      `json:"match_id"`
	TeamID    uint      `json:"team_id"`
	IsHome    bool      `json:"is_home"`
	Formation string    `json:"formation"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
