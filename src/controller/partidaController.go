package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

func CreatePartida(c *gin.Context) {
	var partida model.Partida
	if err := c.ShouldBindJSON(&partida); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&partida).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar partida"})
		return
	}

	c.JSON(http.StatusCreated, partida)
}

func GetPartidas(c *gin.Context) {
	var partidas []model.Partida
	if err := config.DB.Find(&partidas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar partidas"})
		return
	}
	c.JSON(http.StatusOK, partidas)
}

func GetPartidaByID(c *gin.Context) {
	var partida model.Partida
	id := c.Param("id")

	if err := config.DB.First(&partida, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Partida não encontrada"})
		return
	}

	c.JSON(http.StatusOK, partida)
}

func UpdatePartida(c *gin.Context) {
	var partida model.Partida
	id := c.Param("id")

	if err := config.DB.First(&partida, id).Error; err != nil {

	}
}
