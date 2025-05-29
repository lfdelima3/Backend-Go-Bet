package model

import (
	"time"
)

type Tournament struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	Country   string    `json:"country" validate:"required,min=2,max=100"`
	Season    string    `json:"season" validate:"required,min=4,max=9"`
	StartDate time.Time `json:"start_date" validate:"required,future_date"`
	EndDate   time.Time `json:"end_date" validate:"required,future_date"`
	Status    string    `json:"status" validate:"required,oneof=upcoming active finished cancelled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TournamentCreate struct {
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	Country   string    `json:"country" validate:"required,min=2,max=100"`
	Season    string    `json:"season" validate:"required,min=4,max=9"`
	StartDate time.Time `json:"start_date" validate:"required,future_date"`
	EndDate   time.Time `json:"end_date" validate:"required,future_date"`
}

type TournamentUpdate struct {
	Name      string    `json:"name" validate:"omitempty,min=3,max=100"`
	Country   string    `json:"country" validate:"omitempty,min=2,max=100"`
	Season    string    `json:"season" validate:"omitempty,min=4,max=9"`
	StartDate time.Time `json:"start_date" validate:"omitempty,future_date"`
	EndDate   time.Time `json:"end_date" validate:"omitempty,future_date"`
	Status    string    `json:"status" validate:"omitempty,oneof=upcoming active finished cancelled"`
}

type TournamentResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	Season    string    `json:"season"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
