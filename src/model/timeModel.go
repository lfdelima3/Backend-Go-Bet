package model

import (
	"time"
)

type Team struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" validate:"required,min=3,max=100,team_name"`
	Country     string    `json:"country" validate:"required,min=2,max=100"`
	City        string    `json:"city" validate:"required,min=2,max=100"`
	FoundedYear int       `json:"founded_year" validate:"required,min=1800,max=2024"`
	Stadium     string    `json:"stadium" validate:"required,min=3,max=100"`
	Logo        string    `json:"logo" validate:"omitempty,url"`
	Website     string    `json:"website" validate:"omitempty,url"`
	Status      string    `json:"status" validate:"required,oneof=active inactive"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TeamCreate struct {
	Name        string `json:"name" validate:"required,min=3,max=100,team_name"`
	Country     string `json:"country" validate:"required,min=2,max=100"`
	City        string `json:"city" validate:"required,min=2,max=100"`
	FoundedYear int    `json:"founded_year" validate:"required,min=1800,max=2024"`
	Stadium     string `json:"stadium" validate:"required,min=3,max=100"`
	Logo        string `json:"logo" validate:"omitempty,url"`
	Website     string `json:"website" validate:"omitempty,url"`
}

type TeamUpdate struct {
	Name        string `json:"name" validate:"omitempty,min=3,max=100,team_name"`
	Country     string `json:"country" validate:"omitempty,min=2,max=100"`
	City        string `json:"city" validate:"omitempty,min=2,max=100"`
	FoundedYear int    `json:"founded_year" validate:"omitempty,min=1800,max=2024"`
	Stadium     string `json:"stadium" validate:"omitempty,min=3,max=100"`
	Logo        string `json:"logo" validate:"omitempty,url"`
	Website     string `json:"website" validate:"omitempty,url"`
	Status      string `json:"status" validate:"omitempty,oneof=active inactive"`
}

type TeamResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Country     string    `json:"country"`
	City        string    `json:"city"`
	FoundedYear int       `json:"founded_year"`
	Stadium     string    `json:"stadium"`
	Logo        string    `json:"logo"`
	Website     string    `json:"website"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
