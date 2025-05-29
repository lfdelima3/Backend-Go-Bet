package model

import (
	"time"
)

type Bet struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" validate:"required"`
	MatchID   uint      `json:"match_id" validate:"required"`
	BetType   string    `json:"bet_type" validate:"required,oneof=win draw loss over_under corners cards goals"`
	Amount    float64   `json:"amount" validate:"required,min=1,valid_amount"`
	Odds      float64   `json:"odds" validate:"required,min=1,valid_odds"`
	Status    string    `json:"status" validate:"required,oneof=pending won lost cancelled"`
	Result    string    `json:"result" validate:"omitempty,oneof=win draw loss"`
	Payout    float64   `json:"payout" validate:"omitempty,min=0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BetCreate struct {
	UserID  uint    `json:"user_id" validate:"required"`
	MatchID uint    `json:"match_id" validate:"required"`
	BetType string  `json:"bet_type" validate:"required,oneof=win draw loss over_under corners cards goals"`
	Amount  float64 `json:"amount" validate:"required,min=1,valid_amount"`
	Odds    float64 `json:"odds" validate:"required,min=1,valid_odds"`
}

type BetUpdate struct {
	BetType string  `json:"bet_type" validate:"omitempty,oneof=win draw loss over_under corners cards goals"`
	Amount  float64 `json:"amount" validate:"omitempty,min=1,valid_amount"`
	Odds    float64 `json:"odds" validate:"omitempty,min=1,valid_odds"`
	Status  string  `json:"status" validate:"omitempty,oneof=pending won lost cancelled"`
	Result  string  `json:"result" validate:"omitempty,oneof=win draw loss"`
	Payout  float64 `json:"payout" validate:"omitempty,min=0"`
}

type BetResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	MatchID   uint      `json:"match_id"`
	BetType   string    `json:"bet_type"`
	Amount    float64   `json:"amount"`
	Odds      float64   `json:"odds"`
	Status    string    `json:"status"`
	Result    string    `json:"result"`
	Payout    float64   `json:"payout"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
