package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

func CreateSubstituicao(c *gin.Context) {
	var sub model.Substituicao
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&sub).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar substituição"})
		return
	}
	c.JSON(http.StatusCreated, sub)
}

func GetSubstituicoes(c *gin.Context) {
	var subs []model.Substituicao
	if err := config.DB.Find(&subs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar substituições"})
		return
	}
	c.JSON(http.StatusOK, subs)
}

func GetSubstituicaoByID(c *gin.Context) {
	var sub model.Substituicao
	id := c.Param("id")

	if err := config.DB.First(&sub, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Substituição não encontrada"})
		return
	}

	c.JSON(http.StatusOK, sub)
}

func UpdateSubstituicao(c *gin.Context) {
	var sub model.Substituicao
	id := c.Param("id")

	if err := config.DB.First(&sub, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Substituição não encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&sub).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar substituição"})
		return
	}

	c.JSON(http.StatusOK, sub)
}

func DeleteSubstituicao(c *gin.Context) {
	var sub model.Substituicao
	id := c.Param("id")

	if err := config.DB.First(&sub, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Substituição não encontrada"})
		return
	}

	if err := config.DB.Delete(&sub).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar substituição"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Substituição deletada com sucesso"})
}
