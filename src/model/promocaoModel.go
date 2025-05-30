package model

import (
	"time"
)

type Promotion struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"required,min=10,max=500"`
	Type        string    `json:"type" validate:"required,oneof=welcome deposit bet_bonus cashback loyalty"`
	Value       float64   `json:"value" validate:"required,min=0"`
	MinBet      float64   `json:"min_bet" validate:"required,min=0"`
	MaxBet      float64   `json:"max_bet" validate:"required,min=0"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PromotionCreate struct {
	Name        string    `json:"name" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"required,min=10,max=500"`
	Type        string    `json:"type" validate:"required,oneof=welcome deposit bet_bonus cashback loyalty"`
	Value       float64   `json:"value" validate:"required,min=0"`
	MinBet      float64   `json:"min_bet" validate:"required,min=0"`
	MaxBet      float64   `json:"max_bet" validate:"required,min=0"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
}

type PromotionUpdate struct {
	Name        string    `json:"name" validate:"omitempty,min=3,max=100"`
	Description string    `json:"description" validate:"omitempty,min=10,max=500"`
	Type        string    `json:"type" validate:"omitempty,oneof=welcome deposit bet_bonus cashback loyalty"`
	Value       float64   `json:"value" validate:"omitempty,min=0"`
	MinBet      float64   `json:"min_bet" validate:"omitempty,min=0"`
	MaxBet      float64   `json:"max_bet" validate:"omitempty,min=0"`
	StartDate   time.Time `json:"start_date" validate:"omitempty"`
	EndDate     time.Time `json:"end_date" validate:"omitempty"`
	IsActive    bool      `json:"is_active"`
}

type PromotionResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Value       float64   `json:"value"`
	MinBet      float64   `json:"min_bet"`
	MaxBet      float64   `json:"max_bet"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Tipos de promoções
const (
	PromotionTypeWelcome  = "welcome"
	PromotionTypeDeposit  = "deposit"
	PromotionTypeBetBonus = "bet_bonus"
	PromotionTypeCashback = "cashback"
	PromotionTypeLoyalty  = "loyalty"
)
