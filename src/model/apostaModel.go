package model

import (
	"time"
)

type Bet struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       uint      `json:"user_id" validate:"required"`
	MatchID      uint      `json:"match_id" validate:"required"`
	BetType      string    `json:"bet_type" validate:"required,oneof=win draw loss over_under corners cards goals first_goal exact_score asian_handicap both_teams_score total_goals"`
	Amount       float64   `json:"amount" validate:"required,min=1,valid_amount"`
	Odds         float64   `json:"odds" validate:"required,min=1,valid_odds"`
	Status       string    `json:"status" validate:"required,oneof=pending won lost cancelled"`
	Result       string    `json:"result" validate:"omitempty,oneof=win draw loss"`
	Payout       float64   `json:"payout" validate:"omitempty,min=0"`
	CashoutValue float64   `json:"cashout_value" validate:"omitempty,min=0"`
	IsCashout    bool      `json:"is_cashout"`
	BonusApplied float64   `json:"bonus_applied" validate:"omitempty,min=0"`
	PromotionID  *uint     `json:"promotion_id"`
	BetLimit     float64   `json:"bet_limit" validate:"omitempty,min=0"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type BetCreate struct {
	UserID      uint    `json:"user_id" validate:"required"`
	MatchID     uint    `json:"match_id" validate:"required"`
	BetType     string  `json:"bet_type" validate:"required,oneof=win draw loss over_under corners cards goals first_goal exact_score asian_handicap both_teams_score total_goals"`
	Amount      float64 `json:"amount" validate:"required,min=1,valid_amount"`
	Odds        float64 `json:"odds" validate:"required,min=1,valid_odds"`
	PromotionID *uint   `json:"promotion_id"`
}

type BetUpdate struct {
	BetType      string  `json:"bet_type" validate:"omitempty,oneof=win draw loss over_under corners cards goals first_goal exact_score asian_handicap both_teams_score total_goals"`
	Amount       float64 `json:"amount" validate:"omitempty,min=1,valid_amount"`
	Odds         float64 `json:"odds" validate:"omitempty,min=1,valid_odds"`
	Status       string  `json:"status" validate:"omitempty,oneof=pending won lost cancelled"`
	Result       string  `json:"result" validate:"omitempty,oneof=win draw loss"`
	Payout       float64 `json:"payout" validate:"omitempty,min=0"`
	CashoutValue float64 `json:"cashout_value" validate:"omitempty,min=0"`
	IsCashout    bool    `json:"is_cashout"`
	BonusApplied float64 `json:"bonus_applied" validate:"omitempty,min=0"`
}

type BetResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	MatchID      uint      `json:"match_id"`
	BetType      string    `json:"bet_type"`
	Amount       float64   `json:"amount"`
	Odds         float64   `json:"odds"`
	Status       string    `json:"status"`
	Result       string    `json:"result"`
	Payout       float64   `json:"payout"`
	CashoutValue float64   `json:"cashout_value"`
	IsCashout    bool      `json:"is_cashout"`
	BonusApplied float64   `json:"bonus_applied"`
	PromotionID  *uint     `json:"promotion_id"`
	BetLimit     float64   `json:"bet_limit"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Tipos de apostas disponíveis
const (
	BetTypeWin            = "win"
	BetTypeDraw           = "draw"
	BetTypeLoss           = "loss"
	BetTypeOverUnder      = "over_under"
	BetTypeCorners        = "corners"
	BetTypeCards          = "cards"
	BetTypeGoals          = "goals"
	BetTypeFirstGoal      = "first_goal"
	BetTypeExactScore     = "exact_score"
	BetTypeAsianHandicap  = "asian_handicap"
	BetTypeBothTeamsScore = "both_teams_score"
	BetTypeTotalGoals     = "total_goals"
)

// Status das apostas
const (
	BetStatusPending   = "pending"
	BetStatusWon       = "won"
	BetStatusLost      = "lost"
	BetStatusCancelled = "cancelled"
)
