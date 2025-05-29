package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

func CreateAposta(c *gin.Context) {
	var aposta model.Aposta
	usuarioID := c.GetUint("user_id")

	if err := c.ShouldBindJSON(&aposta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aposta.FKUsuarioID = usuarioID
	aposta.DataAposta = time.Now()

	if err := config.DB.Create(&aposta).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar aposta"})
		return
	}

	c.JSON(http.StatusCreated, aposta)
}

func GetApostasUsuario(c *gin.Context) {
	var apostas []model.Aposta
	usuarioID := c.GetUint("user_id")

	if err := config.DB.Where("fk_usuario_id = ?", usuarioID).Find(&apostas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar apostas"})
		return
	}

	c.JSON(http.StatusOK, apostas)
}

func DeleteAposta(c *gin.Context) {
	id := c.Param("id")
	usuarioID := c.GetUint("user_id")

	var aposta model.Aposta
	if err := config.DB.First(&aposta, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aposta não encontrada"})
		return
	}

	if aposta.FKUsuarioID != usuarioID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Você não pode deletar esta aposta"})
		return
	}

	if err := config.DB.Delete(&aposta).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar aposta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aposta deletada com sucesso"})
}
