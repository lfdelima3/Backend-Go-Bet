package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
	"gorm.io/gorm"
)

type TournamentController struct {
	DB *gorm.DB
}

func NewTournamentController(db *gorm.DB) *TournamentController {
	return &TournamentController{DB: db}
}

func (c *TournamentController) CreateTournament(ctx *gin.Context) {
	var tournament model.Tournament
	if err := ctx.ShouldBindJSON(&tournament); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := tournament.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar se já existe um torneio com o mesmo nome
	var existingTournament model.Tournament
	if err := c.DB.Where("name = ?", tournament.Name).First(&existingTournament).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Já existe um torneio com este nome"})
		return
	}

	// Verificar se a data de início é anterior à data de término
	if tournament.StartDate.After(tournament.EndDate) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "A data de início deve ser anterior à data de término"})
		return
	}

	// Definir status inicial como pending
	tournament.Status = "pending"

	if err := c.DB.Create(&tournament).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar torneio"})
		return
	}

	ctx.JSON(http.StatusCreated, tournament)
}

func (c *TournamentController) ListTournaments(ctx *gin.Context) {
	var tournaments []model.Tournament
	query := c.DB.Model(&model.Tournament{})

	// Filtros
	if status := ctx.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}
	if startDate := ctx.Query("start_date"); startDate != "" {
		date, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			query = query.Where("start_date >= ?", date)
		}
	}
	if endDate := ctx.Query("end_date"); endDate != "" {
		date, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			query = query.Where("end_date <= ?", date)
		}
	}

	// Paginação
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var total int64
	query.Count(&total)

	if err := query.Offset(offset).Limit(limit).Find(&tournaments).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar torneios"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": tournaments,
		"meta": gin.H{
			"total":  total,
			"page":   page,
			"limit":  limit,
			"offset": offset,
		},
	})
}

func (c *TournamentController) GetTournament(ctx *gin.Context) {
	id := ctx.Param("id")
	var tournament model.Tournament

	if err := c.DB.First(&tournament, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Torneio não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar torneio"})
		return
	}

	ctx.JSON(http.StatusOK, tournament)
}

func (c *TournamentController) UpdateTournament(ctx *gin.Context) {
	id := ctx.Param("id")
	var tournament model.Tournament
	if err := c.DB.First(&tournament, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Torneio não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar torneio"})
		return
	}

	var updateData model.Tournament
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Verificar se já existe outro torneio com o mesmo nome
	if updateData.Name != "" && updateData.Name != tournament.Name {
		var existingTournament model.Tournament
		if err := c.DB.Where("name = ? AND id != ?", updateData.Name, id).First(&existingTournament).Error; err == nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Já existe um torneio com este nome"})
			return
		}
	}

	// Verificar se a data de início é anterior à data de término
	if !updateData.StartDate.IsZero() && !updateData.EndDate.IsZero() && updateData.StartDate.After(updateData.EndDate) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "A data de início deve ser anterior à data de término"})
		return
	}

	// Atualizar campos
	if updateData.Name != "" {
		tournament.Name = updateData.Name
	}
	if updateData.Description != "" {
		tournament.Description = updateData.Description
	}
	if !updateData.StartDate.IsZero() {
		tournament.StartDate = updateData.StartDate
	}
	if !updateData.EndDate.IsZero() {
		tournament.EndDate = updateData.EndDate
	}
	if updateData.Status != "" {
		tournament.Status = updateData.Status
	}

	if err := c.DB.Save(&tournament).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar torneio"})
		return
	}

	ctx.JSON(http.StatusOK, tournament)
}

func (c *TournamentController) DeleteTournament(ctx *gin.Context) {
	id := ctx.Param("id")
	var tournament model.Tournament

	if err := c.DB.First(&tournament, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Torneio não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar torneio"})
		return
	}

	// Verificar se o torneio tem partidas agendadas
	var matchCount int64
	if err := c.DB.Model(&model.Match{}).Where("tournament_id = ?", id).Count(&matchCount).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar partidas do torneio"})
		return
	}

	if matchCount > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Não é possível excluir um torneio com partidas agendadas"})
		return
	}

	if err := c.DB.Delete(&tournament).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir torneio"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Torneio excluído com sucesso"})
}
