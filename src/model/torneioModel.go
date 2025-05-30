package model

import (
	"time"

	"github.com/lfdelima3/Backend-Go-Bet/src/util"
	"gorm.io/gorm"
)

// Tournament representa um torneio ou campeonato
type Tournament struct {
	gorm.Model
	Name        string    `json:"name" gorm:"not null" validate:"required,min=3,max=100"`
	Description string    `json:"description" gorm:"type:text"`
	StartDate   time.Time `json:"start_date" gorm:"not null" validate:"required,future_date"`
	EndDate     time.Time `json:"end_date" gorm:"not null" validate:"required,future_date"`
	Status      string    `json:"status" gorm:"type:varchar(20);not null;default:'pending'" validate:"required,oneof=pending active completed cancelled"`
	Teams       []Team    `json:"teams" gorm:"many2many:tournament_teams;"`
	Matches     []Match   `json:"matches" gorm:"foreignKey:TournamentID"`
}

// TableName especifica o nome da tabela no banco de dados
func (Tournament) TableName() string {
	return "tournaments"
}

// Validate valida os campos do torneio
func (t *Tournament) Validate() error {
	return util.ValidateStruct(t)
}

// BeforeCreate é um hook do GORM que é executado antes de criar um novo torneio
func (t *Tournament) BeforeCreate(tx *gorm.DB) error {
	return t.Validate()
}

// BeforeUpdate é um hook do GORM que é executado antes de atualizar um torneio
func (t *Tournament) BeforeUpdate(tx *gorm.DB) error {
	return t.Validate()
}

// TournamentTeam representa a relação entre torneios e times
type TournamentTeam struct {
	TournamentID uint      `gorm:"primaryKey"`
	TeamID       uint      `gorm:"primaryKey"`
	JoinedAt     time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

// TableName especifica o nome da tabela no banco de dados
func (TournamentTeam) TableName() string {
	return "tournament_teams"
}
