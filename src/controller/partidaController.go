package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

// CreateMatch cria uma nova partida
func CreateMatch(c *gin.Context) {
	var matchCreate model.MatchCreate
	if err := c.ShouldBindJSON(&matchCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Verifica se os times existem
	var homeTeam, awayTeam model.Team
	if err := config.DB.First(&homeTeam, matchCreate.HomeTeamID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Time da casa não encontrado"})
		return
	}
	if err := config.DB.First(&awayTeam, matchCreate.AwayTeamID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Time visitante não encontrado"})
		return
	}

	// Verifica se o torneio existe
	var tournament model.Tournament
	if err := config.DB.First(&tournament, matchCreate.TournamentID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Torneio não encontrado"})
		return
	}

	// Cria a partida com os dados validados
	match := model.Match{
		TournamentID: matchCreate.TournamentID,
		HomeTeamID:   matchCreate.HomeTeamID,
		AwayTeamID:   matchCreate.AwayTeamID,
		StartTime:    matchCreate.StartTime,
		EndTime:      matchCreate.EndTime,
		Status:       "scheduled",
		Stadium:      matchCreate.Stadium,
		Referee:      matchCreate.Referee,
		Weather:      matchCreate.Weather,
	}

	if err := config.DB.Create(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar partida"})
		return
	}

	response := model.MatchResponse{
		ID:           match.ID,
		TournamentID: match.TournamentID,
		HomeTeamID:   match.HomeTeamID,
		AwayTeamID:   match.AwayTeamID,
		StartTime:    match.StartTime,
		EndTime:      match.EndTime,
		Status:       match.Status,
		HomeScore:    match.HomeScore,
		AwayScore:    match.AwayScore,
		Stadium:      match.Stadium,
		Referee:      match.Referee,
		Attendance:   match.Attendance,
		Weather:      match.Weather,
		CreatedAt:    match.CreatedAt,
		UpdatedAt:    match.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// ListMatches retorna todas as partidas
func ListMatches(c *gin.Context) {
	var matches []model.Match
	if err := config.DB.Find(&matches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar partidas"})
		return
	}

	// Converte para response
	var response []model.MatchResponse
	for _, match := range matches {
		response = append(response, model.MatchResponse{
			ID:           match.ID,
			TournamentID: match.TournamentID,
			HomeTeamID:   match.HomeTeamID,
			AwayTeamID:   match.AwayTeamID,
			StartTime:    match.StartTime,
			EndTime:      match.EndTime,
			Status:       match.Status,
			HomeScore:    match.HomeScore,
			AwayScore:    match.AwayScore,
			Stadium:      match.Stadium,
			Referee:      match.Referee,
			Attendance:   match.Attendance,
			Weather:      match.Weather,
			CreatedAt:    match.CreatedAt,
			UpdatedAt:    match.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetMatch retorna uma partida específica
func GetMatch(c *gin.Context) {
	id := c.Param("id")
	var match model.Match

	if err := config.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partida não encontrada"})
		return
	}

	response := model.MatchResponse{
		ID:           match.ID,
		TournamentID: match.TournamentID,
		HomeTeamID:   match.HomeTeamID,
		AwayTeamID:   match.AwayTeamID,
		StartTime:    match.StartTime,
		EndTime:      match.EndTime,
		Status:       match.Status,
		HomeScore:    match.HomeScore,
		AwayScore:    match.AwayScore,
		Stadium:      match.Stadium,
		Referee:      match.Referee,
		Attendance:   match.Attendance,
		Weather:      match.Weather,
		CreatedAt:    match.CreatedAt,
		UpdatedAt:    match.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateMatch atualiza uma partida
func UpdateMatch(c *gin.Context) {
	id := c.Param("id")
	var match model.Match
	if err := config.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partida não encontrada"})
		return
	}

	var update model.MatchUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Atualiza apenas os campos fornecidos
	if update.TournamentID != 0 {
		// Verifica se o torneio existe
		var tournament model.Tournament
		if err := config.DB.First(&tournament, update.TournamentID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Torneio não encontrado"})
			return
		}
		match.TournamentID = update.TournamentID
	}
	if update.HomeTeamID != 0 {
		// Verifica se o time existe
		var team model.Team
		if err := config.DB.First(&team, update.HomeTeamID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Time da casa não encontrado"})
			return
		}
		match.HomeTeamID = update.HomeTeamID
	}
	if update.AwayTeamID != 0 {
		// Verifica se o time existe
		var team model.Team
		if err := config.DB.First(&team, update.AwayTeamID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Time visitante não encontrado"})
			return
		}
		match.AwayTeamID = update.AwayTeamID
	}
	if !update.StartTime.IsZero() {
		match.StartTime = update.StartTime
	}
	if !update.EndTime.IsZero() {
		match.EndTime = update.EndTime
	}
	if update.Status != "" {
		match.Status = update.Status
	}
	if update.HomeScore >= 0 {
		match.HomeScore = update.HomeScore
	}
	if update.AwayScore >= 0 {
		match.AwayScore = update.AwayScore
	}
	if update.Stadium != "" {
		match.Stadium = update.Stadium
	}
	if update.Referee != "" {
		match.Referee = update.Referee
	}
	if update.Attendance > 0 {
		match.Attendance = update.Attendance
	}
	if update.Weather != "" {
		match.Weather = update.Weather
	}

	if err := config.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar partida"})
		return
	}

	response := model.MatchResponse{
		ID:           match.ID,
		TournamentID: match.TournamentID,
		HomeTeamID:   match.HomeTeamID,
		AwayTeamID:   match.AwayTeamID,
		StartTime:    match.StartTime,
		EndTime:      match.EndTime,
		Status:       match.Status,
		HomeScore:    match.HomeScore,
		AwayScore:    match.AwayScore,
		Stadium:      match.Stadium,
		Referee:      match.Referee,
		Attendance:   match.Attendance,
		Weather:      match.Weather,
		CreatedAt:    match.CreatedAt,
		UpdatedAt:    match.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteMatch remove uma partida
func DeleteMatch(c *gin.Context) {
	id := c.Param("id")
	var match model.Match
	if err := config.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partida não encontrada"})
		return
	}

	// Soft delete - apenas marca como cancelada
	match.Status = "cancelled"
	if err := config.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cancelar partida"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Partida cancelada com sucesso"})
}
