package model

import "gorm.io/gorm"

type PartidaClube struct {
	gorm.Model
	FKClubeID   uint `json:"fk_clube_id"`
	FKPartidaID uint `json:"fk_partida_id"`
}
