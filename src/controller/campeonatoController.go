package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

func CreateCampeonato(c *gin.Context) {
	var campeonato model.Campeonato
	if err := c.ShouldBindJSON(&campeonato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&campeonato).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar campeonato"})
		return
	}
	c.JSON(http.StatusCreated, campeonato)
}

func GetCampeonatos(c *gin.Context) {
	var campeonatos []model.Campeonato
	if err := config.DB.Find(&campeonatos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar campeonatos"})
		return
	}
	c.JSON(http.StatusOK, campeonatos)
}

func GetCampeonatosByID(c *gin.Context) {
	var campeonato model.Campeonato
	id := c.Param("id")

	if err := config.DB.First(&campeonato, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campeonato não encontrado"})
		return
	}

	c.JSON(http.StatusOK, campeonato)
}

func UpdateCampeonato(c *gin.Context) {
	var campeonato model.Campeonato
	id := c.Param("id")

	if err := config.DB.First(&campeonato, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campeonato não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&campeonato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&campeonato).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar campeonato"})
		return
	}

	c.JSON(http.StatusOK, campeonato)
}

func DeleteCampeonato(c *gin.Context) {
	var campeonato model.Campeonato
	id := c.Param("id")

	if err := config.DB.First(&campeonato, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campeonato não encontrado"})
		return
	}

	if err := config.DB.Delete(&campeonato, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar campeonato"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campeonato deletado com sucesso"})
}
