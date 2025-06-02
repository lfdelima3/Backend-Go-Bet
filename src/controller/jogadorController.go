package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
	"github.com/lfdelima3/Backend-Go-Bet/src/util"
)

// CreatePlayer cria um novo jogador
func CreatePlayer(c *gin.Context) {
	var player model.PlayerCreate
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validação dos dados
	if err := util.ValidateStruct(player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Verificar se o time existe
	var team model.Team
	if err := config.DB.First(&team, player.TeamID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time não encontrado"})
		return
	}

	// Verificar se já existe um jogador com o mesmo número no time
	var existingPlayer model.Player
	if err := config.DB.Where("team_id = ? AND number = ?", player.TeamID, player.Number).First(&existingPlayer).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Já existe um jogador com este número no time"})
		return
	}

	// Criar o jogador
	newPlayer := model.Player{
		TeamID:      player.TeamID,
		Name:        player.Name,
		Number:      player.Number,
		Position:    player.Position,
		Nationality: player.Nationality,
		BirthDate:   player.BirthDate,
		Height:      player.Height,
		Weight:      player.Weight,
		Status:      "active",
	}

	if err := config.DB.Create(&newPlayer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar jogador", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPlayer)
}

// ListPlayers retorna todos os jogadores
func ListPlayers(c *gin.Context) {
	var players []model.Player
	query := config.DB.Model(&model.Player{})

	// Filtros
	if teamID := c.Query("team_id"); teamID != "" {
		query = query.Where("team_id = ?", teamID)
	}
	if position := c.Query("position"); position != "" {
		query = query.Where("position = ?", position)
	}
	if nationality := c.Query("nationality"); nationality != "" {
		query = query.Where("nationality = ?", nationality)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}
	if minHeight := c.Query("min_height"); minHeight != "" {
		query = query.Where("height >= ?", minHeight)
	}
	if maxHeight := c.Query("max_height"); maxHeight != "" {
		query = query.Where("height <= ?", maxHeight)
	}
	if minWeight := c.Query("min_weight"); minWeight != "" {
		query = query.Where("weight >= ?", minWeight)
	}
	if maxWeight := c.Query("max_weight"); maxWeight != "" {
		query = query.Where("weight <= ?", maxWeight)
	}

	// Paginação
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var total int64
	query.Count(&total)

	if err := query.Offset(offset).Limit(limit).Find(&players).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar jogadores", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": players,
		"meta": gin.H{
			"total":  total,
			"page":   page,
			"limit":  limit,
			"offset": offset,
		},
	})
}

// GetPlayer retorna um jogador específico
func GetPlayer(c *gin.Context) {
	var player model.Player
	id := c.Param("id")

	if err := config.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jogador não encontrado"})
		return
	}

	c.JSON(http.StatusOK, player)
}

// UpdatePlayer atualiza um jogador
func UpdatePlayer(c *gin.Context) {
	var player model.Player
	id := c.Param("id")

	if err := config.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jogador não encontrado"})
		return
	}

	var updateData model.PlayerUpdate
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validação dos dados
	if err := util.ValidateStruct(updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Se o número do jogador estiver sendo alterado, verificar se já existe outro jogador com o mesmo número no time
	if updateData.Number != 0 && updateData.Number != player.Number {
		var existingPlayer model.Player
		if err := config.DB.Where("team_id = ? AND number = ? AND id != ?", player.TeamID, updateData.Number, id).First(&existingPlayer).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Já existe um jogador com este número no time"})
			return
		}
	}

	// Se o time estiver sendo alterado, verificar se o novo time existe
	if updateData.TeamID != 0 && updateData.TeamID != player.TeamID {
		var team model.Team
		if err := config.DB.First(&team, updateData.TeamID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Time não encontrado"})
			return
		}
	}

	// Atualizar apenas os campos fornecidos
	if updateData.TeamID != 0 {
		player.TeamID = updateData.TeamID
	}
	if updateData.Name != "" {
		player.Name = updateData.Name
	}
	if updateData.Number != 0 {
		player.Number = updateData.Number
	}
	if updateData.Position != "" {
		player.Position = updateData.Position
	}
	if updateData.Nationality != "" {
		player.Nationality = updateData.Nationality
	}
	if !updateData.BirthDate.IsZero() {
		player.BirthDate = updateData.BirthDate
	}
	if updateData.Height != 0 {
		player.Height = updateData.Height
	}
	if updateData.Weight != 0 {
		player.Weight = updateData.Weight
	}
	if updateData.Status != "" {
		player.Status = updateData.Status
	}

	if err := config.DB.Save(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar jogador", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

// DeletePlayer remove um jogador
func DeletePlayer(c *gin.Context) {
	var player model.Player
	id := c.Param("id")

	if err := config.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jogador não encontrado"})
		return
	}

	// Verificar se o jogador está participando de alguma partida em andamento
	var count int64
	if err := config.DB.Model(&model.MatchEvent{}).Where("player_id = ? OR sub_in_player_id = ? OR sub_out_player_id = ?", id, id, id).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar participação do jogador", "details": err.Error()})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Não é possível deletar um jogador que já participou de partidas"})
		return
	}

	if err := config.DB.Delete(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar jogador", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Jogador removido com sucesso"})
}
