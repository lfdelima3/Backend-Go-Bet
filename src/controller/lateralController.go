package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

func CreateLateral(c *gin.Context) {
	var lateral model.Lateral
	if err := c.ShouldBindJSON(&lateral); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&lateral).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar lateral"})
		return
	}
	c.JSON(http.StatusCreated, lateral)
}

func GetLaterais(c *gin.Context) {
	var laterais []model.Lateral
	if err := config.DB.Find(&laterais).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar laterais"})
		return
	}
	c.JSON(http.StatusOK, laterais)
}

func GetLateralByID(c *gin.Context) {
	var lateral model.Lateral
	id := c.Param("id")

	if err := config.DB.First(&lateral, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lateral não encontrada"})
		return
	}

	c.JSON(http.StatusOK, lateral)
}

func UpdateLateral(c *gin.Context) {
	var lateral model.Lateral
	id := c.Param("id")

	if err := config.DB.First(&lateral, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lateral não encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&lateral); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&lateral).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar lateral"})
		return
	}

	c.JSON(http.StatusOK, lateral)
}

func DeleteLateral(c *gin.Context) {
	var lateral model.Lateral
	id := c.Param("id")

	if err := config.DB.First(&lateral, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lateral não encontrada"})
		return
	}

	if err := config.DB.Delete(&lateral).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar lateral"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lateral deletada com sucesso"})
}
