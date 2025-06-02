package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
	"github.com/lfdelima3/Backend-Go-Bet/src/util"
)

// CreateBet cria uma nova aposta
func CreateBet(c *gin.Context) {
	var betCreate model.BetCreate
	userID := c.GetUint("user_id")

	if err := c.ShouldBindJSON(&betCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validação dos dados
	if err := util.ValidateStruct(betCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Verifica se o usuário existe
	var user model.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	// Verifica se o usuário tem saldo suficiente
	if user.Balance < betCreate.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Saldo insuficiente para realizar a aposta"})
		return
	}

	// Verifica se a partida existe e está disponível para apostas
	var match model.Match
	if err := config.DB.First(&match, betCreate.MatchID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partida não encontrada"})
		return
	}

	if match.Status != "scheduled" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A partida não está disponível para apostas"})
		return
	}

	if match.StartTime.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A partida já começou"})
		return
	}

	// Verifica se o usuário já tem uma aposta ativa para esta partida
	var existingBet model.Bet
	if err := config.DB.Where("user_id = ? AND match_id = ? AND status = ?", userID, betCreate.MatchID, "pending").First(&existingBet).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Você já tem uma aposta ativa para esta partida"})
		return
	}

	// Cria a aposta
	bet := model.Bet{
		UserID:    userID,
		MatchID:   betCreate.MatchID,
		BetType:   betCreate.BetType,
		Amount:    betCreate.Amount,
		Odds:      betCreate.Odds,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	// Inicia uma transação
	tx := config.DB.Begin()

	// Atualiza o saldo do usuário
	user.Balance -= betCreate.Amount
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar saldo do usuário", "details": err.Error()})
		return
	}

	// Cria a aposta
	if err := tx.Create(&bet).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar aposta", "details": err.Error()})
		return
	}

	// Commit da transação
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao finalizar transação", "details": err.Error()})
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
	// Paginação
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Filtros
	status := c.Query("status")
	matchID := c.Query("match_id")
	betType := c.Query("bet_type")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	userID := c.GetUint("user_id")
	query := config.DB.Model(&model.Bet{}).Where("user_id = ?", userID)

	// Aplica filtros
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if matchID != "" {
		query = query.Where("match_id = ?", matchID)
	}
	if betType != "" {
		query = query.Where("bet_type = ?", betType)
	}
	if startDate != "" {
		start, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			query = query.Where("created_at >= ?", start)
		}
	}
	if endDate != "" {
		end, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			query = query.Where("created_at <= ?", end)
		}
	}

	var total int64
	query.Count(&total)

	var bets []model.Bet
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&bets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar apostas", "details": err.Error()})
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

	c.JSON(http.StatusOK, gin.H{
		"data": response,
		"meta": gin.H{
			"total":  total,
			"page":   page,
			"limit":  limit,
			"offset": offset,
		},
	})
}

// GetBet retorna uma aposta específica
func GetBet(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")

	var bet model.Bet
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&bet).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aposta não encontrada"})
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

	c.JSON(http.StatusOK, response)
}

// CancelBet cancela uma aposta
func CancelBet(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")

	var bet model.Bet
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&bet).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aposta não encontrada"})
		return
	}

	// Verifica se a aposta pode ser cancelada
	if bet.Status != "pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Apenas apostas pendentes podem ser canceladas"})
		return
	}

	// Verifica se a partida já começou
	var match model.Match
	if err := config.DB.First(&match, bet.MatchID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partida não encontrada"})
		return
	}

	if match.Status != "scheduled" || match.StartTime.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Não é possível cancelar a aposta, a partida já começou"})
		return
	}

	// Inicia uma transação
	tx := config.DB.Begin()

	// Atualiza o saldo do usuário
	var user model.User
	if err := tx.First(&user, userID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	user.Balance += bet.Amount
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar saldo do usuário", "details": err.Error()})
		return
	}

	// Cancela a aposta
	bet.Status = "cancelled"
	if err := tx.Save(&bet).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cancelar aposta", "details": err.Error()})
		return
	}

	// Commit da transação
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao finalizar transação", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aposta cancelada com sucesso"})
}
