package model

import (
	"time"
)

type Match struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	TournamentID uint      `json:"tournament_id" validate:"required"`
	HomeTeamID   uint      `json:"home_team_id" validate:"required"`
	AwayTeamID   uint      `json:"away_team_id" validate:"required"`
	StartTime    time.Time `json:"start_time" validate:"required,future_date"`
	EndTime      time.Time `json:"end_time" validate:"required,future_date"`
	Status       string    `json:"status" validate:"required,oneof=scheduled live finished cancelled postponed"`
	HomeScore    int       `json:"home_score" validate:"min=0,valid_score"`
	AwayScore    int       `json:"away_score" validate:"min=0,valid_score"`
	Stadium      string    `json:"stadium" validate:"required,min=3,max=100"`
	Referee      string    `json:"referee" validate:"required,min=3,max=100"`
	Attendance   int       `json:"attendance" validate:"min=0"`
	Weather      string    `json:"weather" validate:"omitempty,oneof=sunny cloudy rainy snowy windy"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type MatchCreate struct {
	TournamentID uint      `json:"tournament_id" validate:"required"`
	HomeTeamID   uint      `json:"home_team_id" validate:"required"`
	AwayTeamID   uint      `json:"away_team_id" validate:"required"`
	StartTime    time.Time `json:"start_time" validate:"required,future_date"`
	EndTime      time.Time `json:"end_time" validate:"required,future_date"`
	Stadium      string    `json:"stadium" validate:"required,min=3,max=100"`
	Referee      string    `json:"referee" validate:"required,min=3,max=100"`
	Weather      string    `json:"weather" validate:"omitempty,oneof=sunny cloudy rainy snowy windy"`
}

type MatchUpdate struct {
	TournamentID uint      `json:"tournament_id" validate:"omitempty"`
	HomeTeamID   uint      `json:"home_team_id" validate:"omitempty"`
	AwayTeamID   uint      `json:"away_team_id" validate:"omitempty"`
	StartTime    time.Time `json:"start_time" validate:"omitempty,future_date"`
	EndTime      time.Time `json:"end_time" validate:"omitempty,future_date"`
	Status       string    `json:"status" validate:"omitempty,oneof=scheduled live finished cancelled postponed"`
	HomeScore    int       `json:"home_score" validate:"omitempty,min=0,valid_score"`
	AwayScore    int       `json:"away_score" validate:"omitempty,min=0,valid_score"`
	Stadium      string    `json:"stadium" validate:"omitempty,min=3,max=100"`
	Referee      string    `json:"referee" validate:"omitempty,min=3,max=100"`
	Attendance   int       `json:"attendance" validate:"omitempty,min=0"`
	Weather      string    `json:"weather" validate:"omitempty,oneof=sunny cloudy rainy snowy windy"`
}

type MatchResponse struct {
	ID           uint      `json:"id"`
	TournamentID uint      `json:"tournament_id"`
	HomeTeamID   uint      `json:"home_team_id"`
	AwayTeamID   uint      `json:"away_team_id"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Status       string    `json:"status"`
	HomeScore    int       `json:"home_score"`
	AwayScore    int       `json:"away_score"`
	Stadium      string    `json:"stadium"`
	Referee      string    `json:"referee"`
	Attendance   int       `json:"attendance"`
	Weather      string    `json:"weather"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
