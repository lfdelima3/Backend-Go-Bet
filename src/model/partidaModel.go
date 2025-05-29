package model

import (
	"time"

	"gorm.io/gorm"
)

type Partida struct {
	gorm.Model
	DataHora       time.Time `json:"data_hora"`
	Acrescimo      int       `json:"acrescimo"`
	FKCampeonatoID uint      `json:"fk_campeonato_id"`
}
