package model

import (
	"time"

	"gorm.io/gorm"
)

type Aposta struct {
	gorm.Model
	Valor             float64   `json:"valor"`
	Tipo              string    `json:"tipo"`
	ResultadoApostado string    `json:"resultado_apostado"`
	DataAposta        time.Time `json:"data_aposta"`
	FKUsuarioID       uint      `json:"fk_usuario_id"`
	FKPartidaID       uint      `json:"fk_partida_id"`
}
