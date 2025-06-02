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

// CreateMatch cria uma nova partida
func CreateMatch(c *gin.Context) {
	var matchCreate model.MatchCreate
	if err := c.ShouldBindJSON(&matchCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validação dos dados
	if err := util.ValidateStruct(matchCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validação de datas
	if matchCreate.StartTime.After(matchCreate.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data de início deve ser anterior à data de término"})
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

	// Verifica se os times são diferentes
	if matchCreate.HomeTeamID == matchCreate.AwayTeamID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Times da casa e visitante não podem ser iguais"})
		return
	}

	// Verifica se o torneio existe
	var tournament model.Tournament
	if err := config.DB.First(&tournament, matchCreate.TournamentID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Torneio não encontrado"})
		return
	}

	// Verifica se já existe partida no mesmo horário para algum dos times
	var existingMatch model.Match
	if err := config.DB.Where(
		"(home_team_id = ? OR away_team_id = ?) AND start_time <= ? AND end_time >= ?",
		matchCreate.HomeTeamID, matchCreate.HomeTeamID,
		matchCreate.EndTime, matchCreate.StartTime,
	).First(&existingMatch).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Já existe uma partida agendada para este time no mesmo horário"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar partida", "details": err.Error()})
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
	// Paginação
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Filtros
	status := c.Query("status")
	tournamentID := c.Query("tournament_id")
	teamID := c.Query("team_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	query := config.DB.Model(&model.Match{})

	// Aplica filtros
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if tournamentID != "" {
		query = query.Where("tournament_id = ?", tournamentID)
	}
	if teamID != "" {
		query = query.Where("home_team_id = ? OR away_team_id = ?", teamID, teamID)
	}
	if startDate != "" {
		start, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			query = query.Where("start_time >= ?", start)
		}
	}
	if endDate != "" {
		end, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			query = query.Where("end_time <= ?", end)
		}
	}

	var total int64
	query.Count(&total)

	var matches []model.Match
	if err := query.Offset(offset).Limit(limit).Order("start_time DESC").Find(&matches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar partidas", "details": err.Error()})
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

	// Validação dos dados
	if err := util.ValidateStruct(update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validação de datas
	if !update.StartTime.IsZero() && !update.EndTime.IsZero() && update.StartTime.After(update.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data de início deve ser anterior à data de término"})
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

	if update.HomeTeamID != 0 || update.AwayTeamID != 0 {
		// Verifica se os times existem
		if update.HomeTeamID != 0 {
			var team model.Team
			if err := config.DB.First(&team, update.HomeTeamID).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Time da casa não encontrado"})
				return
			}
			match.HomeTeamID = update.HomeTeamID
		}
		if update.AwayTeamID != 0 {
			var team model.Team
			if err := config.DB.First(&team, update.AwayTeamID).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Time visitante não encontrado"})
				return
			}
			match.AwayTeamID = update.AwayTeamID
		}

		// Verifica se os times são diferentes
		if match.HomeTeamID == match.AwayTeamID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Times da casa e visitante não podem ser iguais"})
			return
		}

		// Verifica conflito de horário
		if !update.StartTime.IsZero() || !update.EndTime.IsZero() {
			startTime := update.StartTime
			endTime := update.EndTime
			if startTime.IsZero() {
				startTime = match.StartTime
			}
			if endTime.IsZero() {
				endTime = match.EndTime
			}

			var existingMatch model.Match
			if err := config.DB.Where(
				"id != ? AND (home_team_id = ? OR away_team_id = ?) AND start_time <= ? AND end_time >= ?",
				match.ID, match.HomeTeamID, match.HomeTeamID,
				endTime, startTime,
			).First(&existingMatch).Error; err == nil {
				c.JSON(http.StatusConflict, gin.H{"error": "Já existe uma partida agendada para este time no mesmo horário"})
				return
			}
		}
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar partida", "details": err.Error()})
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

	// Verifica se a partida já começou
	if match.Status == "live" || match.Status == "finished" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Não é possível cancelar uma partida que já começou ou terminou"})
		return
	}

	// Soft delete - apenas marca como cancelada
	match.Status = "cancelled"
	if err := config.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cancelar partida", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Partida cancelada com sucesso"})
}
