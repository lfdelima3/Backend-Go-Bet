package model

import (
	"time"

	"gorm.io/gorm"
)

type Gol struct {
	gorm.Model
	Tempo            time.Time `json:"tempo"`
	Turno            uint8     `json:"turno"` // 1 ou 2
	Tipo             string    `json:"tipo"`  // ex: "Normal", "Pênalti", "Contra"
	FKPartidaClubeID uint      `json:"fk_partida_clube_id"`
	FKJogadorID      uint      `json:"fk_jogador_id_1"`
}
