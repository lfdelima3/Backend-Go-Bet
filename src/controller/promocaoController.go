package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
	"gorm.io/gorm"
)

type PromotionController struct {
	DB *gorm.DB
}

func NewPromotionController(db *gorm.DB) *PromotionController {
	return &PromotionController{DB: db}
}

func (c *PromotionController) CreatePromotion(ctx *gin.Context) {
	var promotion model.PromotionCreate
	if err := ctx.ShouldBindJSON(&promotion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Validar datas
	if promotion.StartDate.After(promotion.EndDate) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data de início deve ser anterior à data de término"})
		return
	}

	// Validar valores
	if promotion.MinBet > promotion.MaxBet {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Valor mínimo da aposta deve ser menor que o valor máximo"})
		return
	}

	newPromotion := model.Promotion{
		Name:        promotion.Name,
		Description: promotion.Description,
		Type:        promotion.Type,
		Value:       promotion.Value,
		MinBet:      promotion.MinBet,
		MaxBet:      promotion.MaxBet,
		StartDate:   promotion.StartDate,
		EndDate:     promotion.EndDate,
		IsActive:    true,
	}

	if err := c.DB.Create(&newPromotion).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar promoção"})
		return
	}

	ctx.JSON(http.StatusCreated, newPromotion)
}

func (c *PromotionController) ListPromotions(ctx *gin.Context) {
	promotionType := ctx.Query("type")
	isActive := ctx.Query("is_active")
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	query := c.DB.Model(&model.Promotion{})

	if promotionType != "" {
		query = query.Where("type = ?", promotionType)
	}
	if isActive != "" {
		query = query.Where("is_active = ?", isActive == "true")
	}
	if startDate != "" {
		date, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			query = query.Where("start_date >= ?", date)
		}
	}
	if endDate != "" {
		date, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			query = query.Where("end_date <= ?", date)
		}
	}

	var promotions []model.Promotion
	if err := query.Find(&promotions).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar promoções"})
		return
	}

	ctx.JSON(http.StatusOK, promotions)
}

func (c *PromotionController) GetPromotion(ctx *gin.Context) {
	id := ctx.Param("id")
	var promotion model.Promotion

	if err := c.DB.First(&promotion, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Promoção não encontrada"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar promoção"})
		return
	}

	ctx.JSON(http.StatusOK, promotion)
}

func (c *PromotionController) UpdatePromotion(ctx *gin.Context) {
	id := ctx.Param("id")
	var promotion model.Promotion
	if err := c.DB.First(&promotion, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Promoção não encontrada"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar promoção"})
		return
	}

	var updateData model.PromotionUpdate
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Validar datas se ambas forem fornecidas
	if !updateData.StartDate.IsZero() && !updateData.EndDate.IsZero() {
		if updateData.StartDate.After(updateData.EndDate) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data de início deve ser anterior à data de término"})
			return
		}
	}

	// Validar valores se ambos forem fornecidos
	if updateData.MinBet > 0 && updateData.MaxBet > 0 {
		if updateData.MinBet > updateData.MaxBet {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Valor mínimo da aposta deve ser menor que o valor máximo"})
			return
		}
	}

	// Atualizar campos
	if updateData.Name != "" {
		promotion.Name = updateData.Name
	}
	if updateData.Description != "" {
		promotion.Description = updateData.Description
	}
	if updateData.Type != "" {
		promotion.Type = updateData.Type
	}
	if updateData.Value > 0 {
		promotion.Value = updateData.Value
	}
	if updateData.MinBet > 0 {
		promotion.MinBet = updateData.MinBet
	}
	if updateData.MaxBet > 0 {
		promotion.MaxBet = updateData.MaxBet
	}
	if !updateData.StartDate.IsZero() {
		promotion.StartDate = updateData.StartDate
	}
	if !updateData.EndDate.IsZero() {
		promotion.EndDate = updateData.EndDate
	}
	promotion.IsActive = updateData.IsActive

	if err := c.DB.Save(&promotion).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar promoção"})
		return
	}

	ctx.JSON(http.StatusOK, promotion)
}

func (c *PromotionController) DeletePromotion(ctx *gin.Context) {
	id := ctx.Param("id")
	var promotion model.Promotion

	if err := c.DB.First(&promotion, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Promoção não encontrada"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar promoção"})
		return
	}

	if err := c.DB.Delete(&promotion).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir promoção"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Promoção excluída com sucesso"})
}

func (c *PromotionController) GetActivePromotions(ctx *gin.Context) {
	now := time.Now()
	var promotions []model.Promotion

	if err := c.DB.Where("is_active = ? AND start_date <= ? AND end_date >= ?", true, now, now).Find(&promotions).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar promoções ativas"})
		return
	}

	ctx.JSON(http.StatusOK, promotions)
}
