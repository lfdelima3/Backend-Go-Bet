package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
	"github.com/lfdelima3/Backend-Go-Bet/src/util"
)

// CreateTeam cria um novo time
func CreateTeam(c *gin.Context) {
	var team model.TeamCreate
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validação dos dados
	if err := util.ValidateStruct(team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Verificar se já existe um time com o mesmo nome
	var existingTeam model.Team
	if err := config.DB.Where("name = ?", team.Name).First(&existingTeam).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Já existe um time com este nome"})
		return
	}

	// Criar o time
	newTeam := model.Team{
		Name:        team.Name,
		Country:     team.Country,
		City:        team.City,
		FoundedYear: team.FoundedYear,
		Stadium:     team.Stadium,
		Logo:        team.Logo,
		Website:     team.Website,
		Status:      "active",
	}

	if err := config.DB.Create(&newTeam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar time", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTeam)
}

// ListTeams retorna todos os times
func ListTeams(c *gin.Context) {
	var teams []model.Team
	query := config.DB.Model(&model.Team{})

	// Filtros
	if country := c.Query("country"); country != "" {
		query = query.Where("country = ?", country)
	}
	if city := c.Query("city"); city != "" {
		query = query.Where("city = ?", city)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}
	if minYear := c.Query("min_year"); minYear != "" {
		query = query.Where("founded_year >= ?", minYear)
	}
	if maxYear := c.Query("max_year"); maxYear != "" {
		query = query.Where("founded_year <= ?", maxYear)
	}

	// Paginação
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var total int64
	query.Count(&total)

	if err := query.Offset(offset).Limit(limit).Find(&teams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar times", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": teams,
		"meta": gin.H{
			"total":  total,
			"page":   page,
			"limit":  limit,
			"offset": offset,
		},
		})
}

// GetTeam retorna um time específico
func GetTeam(c *gin.Context) {
	var team model.Team
	id := c.Param("id")

	if err := config.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time não encontrado"})
		return
	}

	c.JSON(http.StatusOK, team)
}

// UpdateTeam atualiza um time
func UpdateTeam(c *gin.Context) {
	var team model.Team
	id := c.Param("id")

	if err := config.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time não encontrado"})
		return
	}

	var updateData model.TeamUpdate
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validação dos dados
	if err := util.ValidateStruct(updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Se o nome estiver sendo alterado, verificar se já existe outro time com o mesmo nome
	if updateData.Name != "" && updateData.Name != team.Name {
		var existingTeam model.Team
		if err := config.DB.Where("name = ? AND id != ?", updateData.Name, id).First(&existingTeam).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Já existe um time com este nome"})
			return
		}
	}

	// Atualizar apenas os campos fornecidos
	if updateData.Name != "" {
		team.Name = updateData.Name
	}
	if updateData.Country != "" {
		team.Country = updateData.Country
	}
	if updateData.City != "" {
		team.City = updateData.City
	}
	if updateData.FoundedYear != 0 {
		team.FoundedYear = updateData.FoundedYear
	}
	if updateData.Stadium != "" {
		team.Stadium = updateData.Stadium
	}
	if updateData.Logo != "" {
		team.Logo = updateData.Logo
	}
	if updateData.Website != "" {
		team.Website = updateData.Website
	}
	if updateData.Status != "" {
		team.Status = updateData.Status
	}

	if err := config.DB.Save(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar time", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}

// DeleteTeam remove um time
func DeleteTeam(c *gin.Context) {
	var team model.Team
	id := c.Param("id")

	if err := config.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time não encontrado"})
		return
	}

	// Verificar se o time tem jogadores ativos
	var count int64
	if err := config.DB.Model(&model.Player{}).Where("team_id = ? AND status = ?", id, "active").Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar jogadores do time", "details": err.Error()})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Não é possível deletar um time que possui jogadores ativos"})
		return
	}

	// Verificar se o time tem partidas agendadas
	if err := config.DB.Model(&model.Match{}).Where("(home_team_id = ? OR away_team_id = ?) AND status = ?", id, id, "scheduled").Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar partidas do time", "details": err.Error()})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Não é possível deletar um time que possui partidas agendadas"})
		return
	}

	if err := config.DB.Delete(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar time", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Time removido com sucesso"})
}
