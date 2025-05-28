package model

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Nome         string    `json:"nome"`
	DataFuncacao time.Time `json:"data_fundacao"`
	Cidade       string    `json:"cidade"`
	UF           string    `json:"uf"`
}
