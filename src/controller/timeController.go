package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

// CreateTeam cria um novo time
func CreateTeam(c *gin.Context) {
	var teamCreate model.TeamCreate
	if err := c.ShouldBindJSON(&teamCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Verifica se já existe um time com o mesmo nome
	var existingTeam model.Team
	if err := config.DB.Where("name = ?", teamCreate.Name).First(&existingTeam).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Já existe um time com este nome"})
		return
	}

	// Cria o time com os dados validados
	team := model.Team{
		Name:        teamCreate.Name,
		Country:     teamCreate.Country,
		City:        teamCreate.City,
		FoundedYear: teamCreate.FoundedYear,
		Stadium:     teamCreate.Stadium,
		Logo:        teamCreate.Logo,
		Website:     teamCreate.Website,
		Status:      "active",
	}

	if err := config.DB.Create(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar time"})
		return
	}

	response := model.TeamResponse{
		ID:          team.ID,
		Name:        team.Name,
		Country:     team.Country,
		City:        team.City,
		FoundedYear: team.FoundedYear,
		Stadium:     team.Stadium,
		Logo:        team.Logo,
		Website:     team.Website,
		Status:      team.Status,
		CreatedAt:   team.CreatedAt,
		UpdatedAt:   team.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// ListTeams retorna todos os times
func ListTeams(c *gin.Context) {
	var teams []model.Team
	if err := config.DB.Find(&teams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar times"})
		return
	}

	// Converte para response
	var response []model.TeamResponse
	for _, team := range teams {
		response = append(response, model.TeamResponse{
			ID:          team.ID,
			Name:        team.Name,
			Country:     team.Country,
			City:        team.City,
			FoundedYear: team.FoundedYear,
			Stadium:     team.Stadium,
			Logo:        team.Logo,
			Website:     team.Website,
			Status:      team.Status,
			CreatedAt:   team.CreatedAt,
			UpdatedAt:   team.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetTeam retorna um time específico
func GetTeam(c *gin.Context) {
	id := c.Param("id")
	var team model.Team

	if err := config.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time não encontrado"})
		return
	}

	response := model.TeamResponse{
		ID:          team.ID,
		Name:        team.Name,
		Country:     team.Country,
		City:        team.City,
		FoundedYear: team.FoundedYear,
		Stadium:     team.Stadium,
		Logo:        team.Logo,
		Website:     team.Website,
		Status:      team.Status,
		CreatedAt:   team.CreatedAt,
		UpdatedAt:   team.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateTeam atualiza um time
func UpdateTeam(c *gin.Context) {
	id := c.Param("id")
	var team model.Team
	if err := config.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time não encontrado"})
		return
	}

	var update model.TeamUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Atualiza apenas os campos fornecidos
	if update.Name != "" {
		// Verifica se o novo nome já está em uso
		var existingTeam model.Team
		if err := config.DB.Where("name = ? AND id != ?", update.Name, id).First(&existingTeam).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Já existe um time com este nome"})
			return
		}
		team.Name = update.Name
	}
	if update.Country != "" {
		team.Country = update.Country
	}
	if update.City != "" {
		team.City = update.City
	}
	if update.FoundedYear != 0 {
		team.FoundedYear = update.FoundedYear
	}
	if update.Stadium != "" {
		team.Stadium = update.Stadium
	}
	if update.Logo != "" {
		team.Logo = update.Logo
	}
	if update.Website != "" {
		team.Website = update.Website
	}
	if update.Status != "" {
		team.Status = update.Status
	}

	if err := config.DB.Save(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar time"})
		return
	}

	response := model.TeamResponse{
		ID:          team.ID,
		Name:        team.Name,
		Country:     team.Country,
		City:        team.City,
		FoundedYear: team.FoundedYear,
		Stadium:     team.Stadium,
		Logo:        team.Logo,
		Website:     team.Website,
		Status:      team.Status,
		CreatedAt:   team.CreatedAt,
		UpdatedAt:   team.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteTeam remove um time
func DeleteTeam(c *gin.Context) {
	id := c.Param("id")
	var team model.Team
	if err := config.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time não encontrado"})
		return
	}

	// Soft delete - apenas marca como inativo
	team.Status = "inactive"
	if err := config.DB.Save(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao desativar time"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Time desativado com sucesso"})
}
