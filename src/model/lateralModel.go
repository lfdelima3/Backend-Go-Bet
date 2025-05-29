package model

import (
	"time"

	"gorm.io/gorm"
)

type Lateral struct {
	gorm.Model
	Tempo            time.Time `json:"tempo"`
	Turno            uint8     `json:"turno"`
	FKPartidaClubeID uint      `json:"fk_partida_clube_id"`
	FKJogadorID      uint      `json:"fk_jogador_id"`
}

type ThrowIn struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	MatchID     uint      `json:"match_id" validate:"required"`
	TeamID      uint      `json:"team_id" validate:"required"`
	PlayerID    uint      `json:"player_id" validate:"required"`
	Minute      int       `json:"minute" validate:"required,min=1,max=120"`
	Location    string    `json:"location" validate:"required,oneof=defensive midfield offensive"`
	Result      string    `json:"result" validate:"required,oneof=successful unsuccessful"`
	Description string    `json:"description" validate:"omitempty,max=200"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ThrowInCreate struct {
	MatchID     uint   `json:"match_id" validate:"required"`
	TeamID      uint   `json:"team_id" validate:"required"`
	PlayerID    uint   `json:"player_id" validate:"required"`
	Minute      int    `json:"minute" validate:"required,min=1,max=120"`
	Location    string `json:"location" validate:"required,oneof=defensive midfield offensive"`
	Result      string `json:"result" validate:"required,oneof=successful unsuccessful"`
	Description string `json:"description" validate:"omitempty,max=200"`
}

type ThrowInUpdate struct {
	TeamID      uint   `json:"team_id" validate:"omitempty"`
	PlayerID    uint   `json:"player_id" validate:"omitempty"`
	Minute      int    `json:"minute" validate:"omitempty,min=1,max=120"`
	Location    string `json:"location" validate:"omitempty,oneof=defensive midfield offensive"`
	Result      string `json:"result" validate:"omitempty,oneof=successful unsuccessful"`
	Description string `json:"description" validate:"omitempty,max=200"`
}

type ThrowInResponse struct {
	ID          uint      `json:"id"`
	MatchID     uint      `json:"match_id"`
	TeamID      uint      `json:"team_id"`
	PlayerID    uint      `json:"player_id"`
	Minute      int       `json:"minute"`
	Location    string    `json:"location"`
	Result      string    `json:"result"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
