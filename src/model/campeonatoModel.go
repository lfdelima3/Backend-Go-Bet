package model

import (
	"time"

	"gorm.io/gorm"
)

type Campeonato struct {
	gorm.Model
	Nome       string    `json:"nome"`
	DataInicio time.Time `json:"data_inicio"`
	DataFim    time.Time `json:"data_fim"`
	Regiao     string    `json:"regiao"`
	Descricao  string    `json:"descricao"`
}
