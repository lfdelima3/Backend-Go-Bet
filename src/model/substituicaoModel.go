package model

import (
	"time"
)

type Substitution struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	MatchID     uint      `json:"match_id" validate:"required"`
	TeamID      uint      `json:"team_id" validate:"required"`
	PlayerOutID uint      `json:"player_out_id" validate:"required"`
	PlayerInID  uint      `json:"player_in_id" validate:"required"`
	Minute      int       `json:"minute" validate:"required,min=1,max=120"`
	Reason      string    `json:"reason" validate:"required,oneof=tactical injury performance"`
	Description string    `json:"description" validate:"omitempty,max=200"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SubstitutionCreate struct {
	MatchID     uint   `json:"match_id" validate:"required"`
	TeamID      uint   `json:"team_id" validate:"required"`
	PlayerOutID uint   `json:"player_out_id" validate:"required"`
	PlayerInID  uint   `json:"player_in_id" validate:"required"`
	Minute      int    `json:"minute" validate:"required,min=1,max=120"`
	Reason      string `json:"reason" validate:"required,oneof=tactical injury performance"`
	Description string `json:"description" validate:"omitempty,max=200"`
}

type SubstitutionUpdate struct {
	TeamID      uint   `json:"team_id" validate:"omitempty"`
	PlayerOutID uint   `json:"player_out_id" validate:"omitempty"`
	PlayerInID  uint   `json:"player_in_id" validate:"omitempty"`
	Minute      int    `json:"minute" validate:"omitempty,min=1,max=120"`
	Reason      string `json:"reason" validate:"omitempty,oneof=tactical injury performance"`
	Description string `json:"description" validate:"omitempty,max=200"`
}

type SubstitutionResponse struct {
	ID          uint      `json:"id"`
	MatchID     uint      `json:"match_id"`
	TeamID      uint      `json:"team_id"`
	PlayerOutID uint      `json:"player_out_id"`
	PlayerInID  uint      `json:"player_in_id"`
	Minute      int       `json:"minute"`
	Reason      string    `json:"reason"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
