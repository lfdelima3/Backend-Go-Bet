package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

func CreatePartidaClube(c *gin.Context) {
	var matchTeam model.MatchTeamCreate

	if err := c.ShouldBindJSON(&matchTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Criar a relação partida-clube
	newMatchTeam := model.MatchTeam{
		MatchID:   matchTeam.MatchID,
		TeamID:    matchTeam.TeamID,
		IsHome:    matchTeam.IsHome,
		Formation: matchTeam.Formation,
	}

	if err := config.DB.Create(&newMatchTeam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar relação partida-clube"})
		return
	}

	c.JSON(http.StatusCreated, newMatchTeam)
}

func GetPartidasClubes(c *gin.Context) {
	var matchTeams []model.MatchTeam

	if err := config.DB.Find(&matchTeams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar dados"})
		return
	}

	c.JSON(http.StatusOK, matchTeams)
}

func DeletePartidaClube(c *gin.Context) {
	id := c.Param("id")
	var matchTeam model.MatchTeam

	if err := config.DB.First(&matchTeam, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Relação não encontrada"})
		return
	}

	if err := config.DB.Delete(&matchTeam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Relação deletada com sucesso"})
}
