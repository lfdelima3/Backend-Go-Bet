package model

import (
	"time"
)

type Player struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	TeamID      uint      `json:"team_id" validate:"required"`
	Name        string    `json:"name" validate:"required,min=3,max=100"`
	Number      int       `json:"number" validate:"required,min=1,max=99"`
	Position    string    `json:"position" validate:"required,oneof=goalkeeper defender midfielder forward"`
	Nationality string    `json:"nationality" validate:"required,min=2,max=100"`
	BirthDate   time.Time `json:"birth_date" validate:"required"`
	Height      float64   `json:"height" validate:"required,min=1.5,max=2.5"`
	Weight      float64   `json:"weight" validate:"required,min=40,max=150"`
	Status      string    `json:"status" validate:"required,oneof=active inactive injured suspended"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PlayerCreate struct {
	TeamID      uint      `json:"team_id" validate:"required"`
	Name        string    `json:"name" validate:"required,min=3,max=100"`
	Number      int       `json:"number" validate:"required,min=1,max=99"`
	Position    string    `json:"position" validate:"required,oneof=goalkeeper defender midfielder forward"`
	Nationality string    `json:"nationality" validate:"required,min=2,max=100"`
	BirthDate   time.Time `json:"birth_date" validate:"required"`
	Height      float64   `json:"height" validate:"required,min=1.5,max=2.5"`
	Weight      float64   `json:"weight" validate:"required,min=40,max=150"`
}

type PlayerUpdate struct {
	TeamID      uint      `json:"team_id" validate:"omitempty"`
	Name        string    `json:"name" validate:"omitempty,min=3,max=100"`
	Number      int       `json:"number" validate:"omitempty,min=1,max=99"`
	Position    string    `json:"position" validate:"omitempty,oneof=goalkeeper defender midfielder forward"`
	Nationality string    `json:"nationality" validate:"omitempty,min=2,max=100"`
	BirthDate   time.Time `json:"birth_date" validate:"omitempty"`
	Height      float64   `json:"height" validate:"omitempty,min=1.5,max=2.5"`
	Weight      float64   `json:"weight" validate:"omitempty,min=40,max=150"`
	Status      string    `json:"status" validate:"omitempty,oneof=active inactive injured suspended"`
}

type PlayerResponse struct {
	ID          uint      `json:"id"`
	TeamID      uint      `json:"team_id"`
	Name        string    `json:"name"`
	Number      int       `json:"number"`
	Position    string    `json:"position"`
	Nationality string    `json:"nationality"`
	BirthDate   time.Time `json:"birth_date"`
	Height      float64   `json:"height"`
	Weight      float64   `json:"weight"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
