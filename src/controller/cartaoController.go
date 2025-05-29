package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

// CreateCard cria um novo cartão
func CreateCard(c *gin.Context) {
	var cardCreate model.CardCreate
	if err := c.ShouldBindJSON(&cardCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Valida entidades relacionadas
	if err := config.DB.First(&model.Match{}, cardCreate.MatchID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Partida não encontrada"})
		return
	}
	if err := config.DB.First(&model.Team{}, cardCreate.TeamID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Time não encontrado"})
		return
	}

	card := model.Card{
		MatchID:     cardCreate.MatchID,
		TeamID:      cardCreate.TeamID,
		PlayerID:    cardCreate.PlayerID,
		Minute:      cardCreate.Minute,
		CardType:    cardCreate.CardType,
		Reason:      cardCreate.Reason,
		Description: cardCreate.Description,
	}

	if err := config.DB.Create(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cartão"})
		return
	}

	response := model.CardResponse{
		ID:          card.ID,
		MatchID:     card.MatchID,
		TeamID:      card.TeamID,
		PlayerID:    card.PlayerID,
		Minute:      card.Minute,
		CardType:    card.CardType,
		Reason:      card.Reason,
		Description: card.Description,
		CreatedAt:   card.CreatedAt,
		UpdatedAt:   card.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// ListCards retorna todos os cartões
func ListCards(c *gin.Context) {
	var cards []model.Card
	if err := config.DB.Find(&cards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar cartões"})
		return
	}

	var response []model.CardResponse
	for _, card := range cards {
		response = append(response, model.CardResponse{
			ID:          card.ID,
			MatchID:     card.MatchID,
			TeamID:      card.TeamID,
			PlayerID:    card.PlayerID,
			Minute:      card.Minute,
			CardType:    card.CardType,
			Reason:      card.Reason,
			Description: card.Description,
			CreatedAt:   card.CreatedAt,
			UpdatedAt:   card.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetCard retorna um cartão específico
func GetCard(c *gin.Context) {
	id := c.Param("id")
	var card model.Card
	if err := config.DB.First(&card, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cartão não encontrado"})
		return
	}

	response := model.CardResponse{
		ID:          card.ID,
		MatchID:     card.MatchID,
		TeamID:      card.TeamID,
		PlayerID:    card.PlayerID,
		Minute:      card.Minute,
		CardType:    card.CardType,
		Reason:      card.Reason,
		Description: card.Description,
		CreatedAt:   card.CreatedAt,
		UpdatedAt:   card.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateCard atualiza um cartão
func UpdateCard(c *gin.Context) {
	id := c.Param("id")
	var card model.Card
	if err := config.DB.First(&card, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cartão não encontrado"})
		return
	}

	var update model.CardUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	if update.TeamID != 0 {
		card.TeamID = update.TeamID
	}
	if update.PlayerID != 0 {
		card.PlayerID = update.PlayerID
	}
	if update.Minute > 0 {
		card.Minute = update.Minute
	}
	if update.CardType != "" {
		card.CardType = update.CardType
	}
	if update.Reason != "" {
		card.Reason = update.Reason
	}
	if update.Description != "" {
		card.Description = update.Description
	}

	if err := config.DB.Save(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar cartão"})
		return
	}

	response := model.CardResponse{
		ID:          card.ID,
		MatchID:     card.MatchID,
		TeamID:      card.TeamID,
		PlayerID:    card.PlayerID,
		Minute:      card.Minute,
		CardType:    card.CardType,
		Reason:      card.Reason,
		Description: card.Description,
		CreatedAt:   card.CreatedAt,
		UpdatedAt:   card.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteCard remove um cartão
func DeleteCard(c *gin.Context) {
	id := c.Param("id")
	var card model.Card
	if err := config.DB.First(&card, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cartão não encontrado"})
		return
	}

	if err := config.DB.Delete(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar cartão"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cartão deletado com sucesso"})
}
