package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

// CreateGoal registra um novo gol
func CreateGoal(c *gin.Context) {
	var goalCreate model.GoalCreate
	if err := c.ShouldBindJSON(&goalCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Valida entidades relacionadas
	if err := config.DB.First(&model.Match{}, goalCreate.MatchID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Partida não encontrada"})
		return
	}
	if err := config.DB.First(&model.Team{}, goalCreate.TeamID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Time não encontrado"})
		return
	}
	// Não existe model.Player, então não validar PlayerID

	goal := model.Goal{
		MatchID:     goalCreate.MatchID,
		TeamID:      goalCreate.TeamID,
		PlayerID:    goalCreate.PlayerID,
		Minute:      goalCreate.Minute,
		IsOwnGoal:   goalCreate.IsOwnGoal,
		IsPenalty:   goalCreate.IsPenalty,
		AssistID:    goalCreate.AssistID,
		Description: goalCreate.Description,
	}

	if err := config.DB.Create(&goal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar gol"})
		return
	}

	response := model.GoalResponse{
		ID:          goal.ID,
		MatchID:     goal.MatchID,
		TeamID:      goal.TeamID,
		PlayerID:    goal.PlayerID,
		Minute:      goal.Minute,
		IsOwnGoal:   goal.IsOwnGoal,
		IsPenalty:   goal.IsPenalty,
		AssistID:    goal.AssistID,
		Description: goal.Description,
		CreatedAt:   goal.CreatedAt,
		UpdatedAt:   goal.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// ListGoals retorna todos os gols
func ListGoals(c *gin.Context) {
	var goals []model.Goal
	if err := config.DB.Find(&goals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar gols"})
		return
	}

	var response []model.GoalResponse
	for _, goal := range goals {
		response = append(response, model.GoalResponse{
			ID:          goal.ID,
			MatchID:     goal.MatchID,
			TeamID:      goal.TeamID,
			PlayerID:    goal.PlayerID,
			Minute:      goal.Minute,
			IsOwnGoal:   goal.IsOwnGoal,
			IsPenalty:   goal.IsPenalty,
			AssistID:    goal.AssistID,
			Description: goal.Description,
			CreatedAt:   goal.CreatedAt,
			UpdatedAt:   goal.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetGoal retorna um gol específico
func GetGoal(c *gin.Context) {
	id := c.Param("id")
	var goal model.Goal

	if err := config.DB.First(&goal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gol não encontrado"})
		return
	}

	response := model.GoalResponse{
		ID:          goal.ID,
		MatchID:     goal.MatchID,
		TeamID:      goal.TeamID,
		PlayerID:    goal.PlayerID,
		Minute:      goal.Minute,
		IsOwnGoal:   goal.IsOwnGoal,
		IsPenalty:   goal.IsPenalty,
		AssistID:    goal.AssistID,
		Description: goal.Description,
		CreatedAt:   goal.CreatedAt,
		UpdatedAt:   goal.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateGoal atualiza um gol
func UpdateGoal(c *gin.Context) {
	id := c.Param("id")
	var goal model.Goal
	if err := config.DB.First(&goal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gol não encontrado"})
		return
	}

	var update model.GoalUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	if update.TeamID != 0 {
		goal.TeamID = update.TeamID
	}
	if update.PlayerID != 0 {
		goal.PlayerID = update.PlayerID
	}
	if update.Minute > 0 {
		goal.Minute = update.Minute
	}
	goal.IsOwnGoal = update.IsOwnGoal
	goal.IsPenalty = update.IsPenalty
	if update.AssistID != 0 {
		goal.AssistID = update.AssistID
	}
	if update.Description != "" {
		goal.Description = update.Description
	}

	if err := config.DB.Save(&goal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar gol"})
		return
	}

	response := model.GoalResponse{
		ID:          goal.ID,
		MatchID:     goal.MatchID,
		TeamID:      goal.TeamID,
		PlayerID:    goal.PlayerID,
		Minute:      goal.Minute,
		IsOwnGoal:   goal.IsOwnGoal,
		IsPenalty:   goal.IsPenalty,
		AssistID:    goal.AssistID,
		Description: goal.Description,
		CreatedAt:   goal.CreatedAt,
		UpdatedAt:   goal.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteGoal remove um gol
func DeleteGoal(c *gin.Context) {
	id := c.Param("id")
	var goal model.Goal
	if err := config.DB.First(&goal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gol não encontrado"})
		return
	}

	if err := config.DB.Delete(&goal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar gol"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gol deletado com sucesso"})
}
