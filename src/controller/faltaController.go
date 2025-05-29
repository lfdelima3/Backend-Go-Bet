package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
	"github.com/lfdelima3/Backend-Go-Bet/src/util"
)

// CreateFoul registra uma nova falta
func CreateFoul(c *gin.Context) {
	var foulCreate model.FoulCreate
	if err := c.ShouldBindJSON(&foulCreate); err != nil {
		util.LogError("Erro ao validar dados da falta", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	// Valida entidades relacionadas
	if err := config.DB.First(&model.Match{}, foulCreate.MatchID).Error; err != nil {
		util.LogError(fmt.Sprintf("Partida não encontrada: %d", foulCreate.MatchID), err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Partida não encontrada"})
		return
	}

	if err := config.DB.First(&model.Team{}, foulCreate.TeamID).Error; err != nil {
		util.LogError(fmt.Sprintf("Time não encontrado: %d", foulCreate.TeamID), err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Time não encontrado"})
		return
	}

	foul := model.Foul{
		MatchID:     foulCreate.MatchID,
		TeamID:      foulCreate.TeamID,
		PlayerID:    foulCreate.PlayerID,
		Minute:      foulCreate.Minute,
		FoulType:    foulCreate.FoulType,
		Location:    foulCreate.Location,
		Description: foulCreate.Description,
	}

	if err := config.DB.Create(&foul).Error; err != nil {
		util.LogError("Erro ao registrar falta", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar falta"})
		return
	}

	response := model.FoulResponse{
		ID:          foul.ID,
		MatchID:     foul.MatchID,
		TeamID:      foul.TeamID,
		PlayerID:    foul.PlayerID,
		Minute:      foul.Minute,
		FoulType:    foul.FoulType,
		Location:    foul.Location,
		Description: foul.Description,
		CreatedAt:   foul.CreatedAt,
		UpdatedAt:   foul.UpdatedAt,
	}

	util.LogInfo(fmt.Sprintf("Falta registrada com sucesso: ID=%d", foul.ID))
	c.JSON(http.StatusCreated, response)
}

// ListFouls retorna todas as faltas
func ListFouls(c *gin.Context) {
	var fouls []model.Foul
	if err := config.DB.Find(&fouls).Error; err != nil {
		util.LogError("Erro ao buscar faltas", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar faltas"})
		return
	}

	var response []model.FoulResponse
	for _, foul := range fouls {
		response = append(response, model.FoulResponse{
			ID:          foul.ID,
			MatchID:     foul.MatchID,
			TeamID:      foul.TeamID,
			PlayerID:    foul.PlayerID,
			Minute:      foul.Minute,
			FoulType:    foul.FoulType,
			Location:    foul.Location,
			Description: foul.Description,
			CreatedAt:   foul.CreatedAt,
			UpdatedAt:   foul.UpdatedAt,
		})
	}

	util.LogInfo(fmt.Sprintf("Listadas %d faltas", len(response)))
	c.JSON(http.StatusOK, response)
}

// GetFoul retorna uma falta específica
func GetFoul(c *gin.Context) {
	id := c.Param("id")
	var foul model.Foul

	// Tenta buscar do cache primeiro
	cacheKey := fmt.Sprintf("foul:%s", id)
	if cachedFoul, err := util.GetCache(cacheKey); err == nil {
		c.JSON(http.StatusOK, cachedFoul)
		return
	}

	if err := config.DB.First(&foul, id).Error; err != nil {
		util.LogError(fmt.Sprintf("Falta não encontrada: %s", id), err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Falta não encontrada"})
		return
	}

	response := model.FoulResponse{
		ID:          foul.ID,
		MatchID:     foul.MatchID,
		TeamID:      foul.TeamID,
		PlayerID:    foul.PlayerID,
		Minute:      foul.Minute,
		FoulType:    foul.FoulType,
		Location:    foul.Location,
		Description: foul.Description,
		CreatedAt:   foul.CreatedAt,
		UpdatedAt:   foul.UpdatedAt,
	}

	// Salva no cache por 5 minutos
	util.SetCache(cacheKey, response, 5*time.Minute)

	util.LogInfo(fmt.Sprintf("Falta recuperada: ID=%s", id))
	c.JSON(http.StatusOK, response)
}

// UpdateFoul atualiza uma falta
func UpdateFoul(c *gin.Context) {
	id := c.Param("id")
	var foul model.Foul
	if err := config.DB.First(&foul, id).Error; err != nil {
		util.LogError(fmt.Sprintf("Falta não encontrada: %s", id), err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Falta não encontrada"})
		return
	}

	var update model.FoulUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		util.LogError("Erro ao validar dados de atualização", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	// Atualiza apenas os campos fornecidos
	if update.TeamID != 0 {
		foul.TeamID = update.TeamID
	}
	if update.PlayerID != 0 {
		foul.PlayerID = update.PlayerID
	}
	if update.Minute > 0 {
		foul.Minute = update.Minute
	}
	if update.FoulType != "" {
		foul.FoulType = update.FoulType
	}
	if update.Location != "" {
		foul.Location = update.Location
	}
	if update.Description != "" {
		foul.Description = update.Description
	}

	if err := config.DB.Save(&foul).Error; err != nil {
		util.LogError("Erro ao atualizar falta", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar falta"})
		return
	}

	// Remove do cache
	cacheKey := fmt.Sprintf("foul:%s", id)
	util.DeleteCache(cacheKey)

	response := model.FoulResponse{
		ID:          foul.ID,
		MatchID:     foul.MatchID,
		TeamID:      foul.TeamID,
		PlayerID:    foul.PlayerID,
		Minute:      foul.Minute,
		FoulType:    foul.FoulType,
		Location:    foul.Location,
		Description: foul.Description,
		CreatedAt:   foul.CreatedAt,
		UpdatedAt:   foul.UpdatedAt,
	}

	util.LogInfo(fmt.Sprintf("Falta atualizada: ID=%s", id))
	c.JSON(http.StatusOK, response)
}

// DeleteFoul remove uma falta
func DeleteFoul(c *gin.Context) {
	id := c.Param("id")
	var foul model.Foul
	if err := config.DB.First(&foul, id).Error; err != nil {
		util.LogError(fmt.Sprintf("Falta não encontrada: %s", id), err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Falta não encontrada"})
		return
	}

	if err := config.DB.Delete(&foul).Error; err != nil {
		util.LogError("Erro ao deletar falta", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar falta"})
		return
	}

	// Remove do cache
	cacheKey := fmt.Sprintf("foul:%s", id)
	util.DeleteCache(cacheKey)

	util.LogInfo(fmt.Sprintf("Falta deletada: ID=%s", id))
	c.JSON(http.StatusOK, gin.H{"message": "Falta deletada com sucesso"})
}
