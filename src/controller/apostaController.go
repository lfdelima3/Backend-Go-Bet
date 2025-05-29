package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

// CreateBet cria uma nova aposta
func CreateBet(c *gin.Context) {
	var betCreate model.BetCreate
	userID := c.GetUint("user_id")

	if err := c.ShouldBindJSON(&betCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	bet := model.Bet{
		UserID:    userID,
		MatchID:   betCreate.MatchID,
		BetType:   betCreate.BetType,
		Amount:    betCreate.Amount,
		Odds:      betCreate.Odds,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	if err := config.DB.Create(&bet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar aposta"})
		return
	}

	response := model.BetResponse{
		ID:        bet.ID,
		UserID:    bet.UserID,
		MatchID:   bet.MatchID,
		BetType:   bet.BetType,
		Amount:    bet.Amount,
		Odds:      bet.Odds,
		Status:    bet.Status,
		Result:    bet.Result,
		Payout:    bet.Payout,
		CreatedAt: bet.CreatedAt,
		UpdatedAt: bet.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// ListUserBets retorna todas as apostas do usuário
func ListUserBets(c *gin.Context) {
	var bets []model.Bet
	userID := c.GetUint("user_id")

	if err := config.DB.Where("user_id = ?", userID).Find(&bets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar apostas"})
		return
	}

	var response []model.BetResponse
	for _, bet := range bets {
		response = append(response, model.BetResponse{
			ID:        bet.ID,
			UserID:    bet.UserID,
			MatchID:   bet.MatchID,
			BetType:   bet.BetType,
			Amount:    bet.Amount,
			Odds:      bet.Odds,
			Status:    bet.Status,
			Result:    bet.Result,
			Payout:    bet.Payout,
			CreatedAt: bet.CreatedAt,
			UpdatedAt: bet.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}

// DeleteBet remove uma aposta do usuário
func DeleteBet(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")

	var bet model.Bet
	if err := config.DB.First(&bet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aposta não encontrada"})
		return
	}

	if bet.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Você não pode deletar esta aposta"})
		return
	}

	if err := config.DB.Delete(&bet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar aposta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aposta deletada com sucesso"})
}
