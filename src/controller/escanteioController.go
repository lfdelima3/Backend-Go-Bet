package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

func CreateEscanteio(c *gin.Context) {
	var escanteio model.Escanteio
	if err := c.ShouldBindJSON(&escanteio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&escanteio).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar escanteio"})
		return
	}
	c.JSON(http.StatusCreated, escanteio)
}

func GetEscanteios(c *gin.Context) {
	var escanteios []model.Escanteio
	if err := config.DB.Find(&escanteios).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar escanteios"})
		return
	}
	c.JSON(http.StatusOK, escanteios)
}

func GetEscanteioByID(c *gin.Context) {
	var escanteio model.Escanteio
	id := c.Param("id")

	if err := config.DB.First(&escanteio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Escanteio não encontrado"})
		return
	}

	c.JSON(http.StatusOK, escanteio)
}

func UpdateEscanteio(c *gin.Context) {
	var escanteio model.Escanteio
	id := c.Param("id")

	if err := config.DB.First(&escanteio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Escanteio não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&escanteio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&escanteio).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar escanteio"})
		return
	}

	c.JSON(http.StatusOK, escanteio)
}

func DeleteEscanteio(c *gin.Context) {
	var escanteio model.Escanteio
	id := c.Param("id")

	if err := config.DB.First(&escanteio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Escanteio não encontrado"})
		return
	}

	if err := config.DB.Delete(&escanteio).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar escanteio"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Escanteio deletado com sucesso"})
}
