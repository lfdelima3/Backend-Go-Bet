package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

func CreateClube(c *gin.Context) {
	var clube model.Team
	if err := c.ShouldBindJSON(&clube); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&clube).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar o clube"})
		return
	}

	c.JSON(http.StatusCreated, clube)
}

func GetClubes(c *gin.Context) {
	var clubes []model.Team
	if err := config.DB.Find(&clubes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar clubes"})
		return
	}
	c.JSON(http.StatusOK, clubes)
}

func GetClubePorID(c *gin.Context) {
	var clube model.Team
	id := c.Param("id")

	if err := config.DB.First(&clube, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clube não encontrado"})
		return
	}
	c.JSON(http.StatusOK, clube)
}

func UpdateClube(c *gin.Context) {
	var clube model.Team
	id := c.Param("id")

	if err := config.DB.First(&clube, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clube não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&clube); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&clube).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar clube"})
		return
	}
	c.JSON(http.StatusOK, clube)
}

func DeleteClube(c *gin.Context) {
	var clube model.Team
	id := c.Param("id")

	if err := config.DB.First(&clube, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clube não encontrado"})
		return
	}

	if err := config.DB.Delete(&clube); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar clube"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Clube deletado com sucesso"})
}
