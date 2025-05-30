package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
	"gorm.io/gorm"
)

type MatchEventController struct {
	DB *gorm.DB
}

func NewMatchEventController(db *gorm.DB) *MatchEventController {
	return &MatchEventController{DB: db}
}

func (c *MatchEventController) CreateEvent(ctx *gin.Context) {
	var event model.MatchEventCreate
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Validar se a partida existe
	var match model.Match
	if err := c.DB.First(&match, event.MatchID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Partida não encontrada"})
		return
	}

	// Validar se o time existe
	var team model.Team
	if err := c.DB.First(&team, event.TeamID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Time não encontrado"})
		return
	}

	// Validar se o jogador existe
	var player model.Player
	if err := c.DB.First(&player, event.PlayerID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Jogador não encontrado"})
		return
	}

	// Validar campos específicos do tipo de evento
	if err := c.validateEventFields(event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEvent := model.MatchEvent{
		MatchID:        event.MatchID,
		EventType:      event.EventType,
		TeamID:         event.TeamID,
		PlayerID:       event.PlayerID,
		Minute:         event.Minute,
		Description:    event.Description,
		GoalType:       event.GoalType,
		CardType:       event.CardType,
		FoulType:       event.FoulType,
		SubInPlayerID:  event.SubInPlayerID,
		SubOutPlayerID: event.SubOutPlayerID,
		ThrowInSide:    event.ThrowInSide,
		CornerSide:     event.CornerSide,
	}

	if err := c.DB.Create(&newEvent).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar evento"})
		return
	}

	ctx.JSON(http.StatusCreated, newEvent)
}

func (c *MatchEventController) ListEvents(ctx *gin.Context) {
	matchID := ctx.Query("match_id")
	eventType := ctx.Query("event_type")
	teamID := ctx.Query("team_id")
	playerID := ctx.Query("player_id")

	query := c.DB.Model(&model.MatchEvent{})

	if matchID != "" {
		query = query.Where("match_id = ?", matchID)
	}
	if eventType != "" {
		query = query.Where("event_type = ?", eventType)
	}
	if teamID != "" {
		query = query.Where("team_id = ?", teamID)
	}
	if playerID != "" {
		query = query.Where("player_id = ?", playerID)
	}

	var events []model.MatchEvent
	if err := query.Find(&events).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar eventos"})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func (c *MatchEventController) GetEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	var event model.MatchEvent

	if err := c.DB.First(&event, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Evento não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar evento"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (c *MatchEventController) UpdateEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	var event model.MatchEvent
	if err := c.DB.First(&event, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Evento não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar evento"})
		return
	}

	var updateData model.MatchEventUpdate
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Atualizar campos
	if updateData.EventType != "" {
		event.EventType = updateData.EventType
	}
	if updateData.TeamID != 0 {
		event.TeamID = updateData.TeamID
	}
	if updateData.PlayerID != 0 {
		event.PlayerID = updateData.PlayerID
	}
	if updateData.Minute != 0 {
		event.Minute = updateData.Minute
	}
	if updateData.Description != "" {
		event.Description = updateData.Description
	}
	if updateData.GoalType != "" {
		event.GoalType = updateData.GoalType
	}
	if updateData.CardType != "" {
		event.CardType = updateData.CardType
	}
	if updateData.FoulType != "" {
		event.FoulType = updateData.FoulType
	}
	if updateData.SubInPlayerID != 0 {
		event.SubInPlayerID = updateData.SubInPlayerID
	}
	if updateData.SubOutPlayerID != 0 {
		event.SubOutPlayerID = updateData.SubOutPlayerID
	}
	if updateData.ThrowInSide != "" {
		event.ThrowInSide = updateData.ThrowInSide
	}
	if updateData.CornerSide != "" {
		event.CornerSide = updateData.CornerSide
	}

	if err := c.DB.Save(&event).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar evento"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (c *MatchEventController) DeleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	var event model.MatchEvent

	if err := c.DB.First(&event, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Evento não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar evento"})
		return
	}

	if err := c.DB.Delete(&event).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir evento"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Evento excluído com sucesso"})
}

func (c *MatchEventController) validateEventFields(event model.MatchEventCreate) error {
	switch event.EventType {
	case model.EventTypeGoal:
		if event.GoalType == "" {
			return fmt.Errorf("tipo de gol é obrigatório")
		}
	case model.EventTypeCard:
		if event.CardType == "" {
			return fmt.Errorf("tipo de cartão é obrigatório")
		}
	case model.EventTypeFoul:
		if event.FoulType == "" {
			return fmt.Errorf("tipo de falta é obrigatório")
		}
	case model.EventTypeSubstitution:
		if event.SubInPlayerID == 0 || event.SubOutPlayerID == 0 {
			return fmt.Errorf("jogadores de substituição são obrigatórios")
		}
	case model.EventTypeThrowIn:
		if event.ThrowInSide == "" {
			return fmt.Errorf("lado do lateral é obrigatório")
		}
	case model.EventTypeCorner:
		if event.CornerSide == "" {
			return fmt.Errorf("lado do escanteio é obrigatório")
		}
	}
	return nil
}
