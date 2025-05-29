package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

func CreateGol(c *gin.Context) {
	var gol model.Gol
	if err := c.ShouldBindJSON(&gol); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&gol).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar gol"})
		return
	}

	c.JSON(http.StatusCreated, gol)
}

func GetGols(c *gin.Context) {
	var gols []model.Gol
	if err := config.DB.Find(&gols).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar gols"})
		return
	}
	c.JSON(http.StatusOK, gols)
}

func GetGolByID(c *gin.Context) {
	var gol model.Gol
	id := c.Param("id")

	if err := config.DB.First(&gol, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gol não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gol)
}

func UpdateGol(c *gin.Context) {
	var gol model.Gol
	id := c.Param("id")

	if err := config.DB.First(&gol, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gol não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&gol); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&gol).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar gol"})
		return
	}

	c.JSON(http.StatusOK, gol)
}

func DeleteGol(c *gin.Context) {
	var gol model.Gol
	id := c.Param("id")

	if err := config.DB.First(&gol, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gol não encontrado"})
		return
	}

	if err := config.DB.Delete(&gol).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar gol"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gol deletado com sucesso"})
}
