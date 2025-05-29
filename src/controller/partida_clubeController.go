package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

func CreatePartidaClube(c *gin.Context) {
	var pc model.PartidaClube

	if err := c.ShouldBindJSON(&pc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&pc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar relação partida-clube"})
		return
	}

	c.JSON(http.StatusCreated, pc)
}

func GetPartidasClubes(c *gin.Context) {
	var relacoes []model.PartidaClube

	if err := config.DB.Find(&relacoes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar dados"})
		return
	}

	c.JSON(http.StatusOK, relacoes)
}

func DeletePartidaClube(c *gin.Context) {
	id := c.Param("id")
	var pc model.PartidaClube

	if err := config.DB.First(&pc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Relação não encontrada"})
		return
	}

	if err := config.DB.Delete(&pc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Relação deletada com sucesso"})
}
