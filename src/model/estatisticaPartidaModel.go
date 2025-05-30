package model

import (
	"time"
)

type MatchStatistics struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	MatchID           uint      `json:"match_id" validate:"required"`
	HomeTeamID        uint      `json:"home_team_id" validate:"required"`
	AwayTeamID        uint      `json:"away_team_id" validate:"required"`
	HomePossession    float64   `json:"home_possession" validate:"min=0,max=100"`
	AwayPossession    float64   `json:"away_possession" validate:"min=0,max=100"`
	HomeShots         int       `json:"home_shots" validate:"min=0"`
	AwayShots         int       `json:"away_shots" validate:"min=0"`
	HomeShotsOnTarget int       `json:"home_shots_on_target" validate:"min=0"`
	AwayShotsOnTarget int       `json:"away_shots_on_target" validate:"min=0"`
	HomeCorners       int       `json:"home_corners" validate:"min=0"`
	AwayCorners       int       `json:"away_corners" validate:"min=0"`
	HomeFouls         int       `json:"home_fouls" validate:"min=0"`
	AwayFouls         int       `json:"away_fouls" validate:"min=0"`
	HomeYellowCards   int       `json:"home_yellow_cards" validate:"min=0"`
	AwayYellowCards   int       `json:"away_yellow_cards" validate:"min=0"`
	HomeRedCards      int       `json:"home_red_cards" validate:"min=0"`
	AwayRedCards      int       `json:"away_red_cards" validate:"min=0"`
	HomeOffsides      int       `json:"home_offsides" validate:"min=0"`
	AwayOffsides      int       `json:"away_offsides" validate:"min=0"`
	HomePasses        int       `json:"home_passes" validate:"min=0"`
	AwayPasses        int       `json:"away_passes" validate:"min=0"`
	HomePassAccuracy  float64   `json:"home_pass_accuracy" validate:"min=0,max=100"`
	AwayPassAccuracy  float64   `json:"away_pass_accuracy" validate:"min=0,max=100"`
	HomeSaves         int       `json:"home_saves" validate:"min=0"`
	AwaySaves         int       `json:"away_saves" validate:"min=0"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type MatchStatisticsCreate struct {
	MatchID           uint    `json:"match_id" validate:"required"`
	HomeTeamID        uint    `json:"home_team_id" validate:"required"`
	AwayTeamID        uint    `json:"away_team_id" validate:"required"`
	HomePossession    float64 `json:"home_possession" validate:"min=0,max=100"`
	AwayPossession    float64 `json:"away_possession" validate:"min=0,max=100"`
	HomeShots         int     `json:"home_shots" validate:"min=0"`
	AwayShots         int     `json:"away_shots" validate:"min=0"`
	HomeShotsOnTarget int     `json:"home_shots_on_target" validate:"min=0"`
	AwayShotsOnTarget int     `json:"away_shots_on_target" validate:"min=0"`
	HomeCorners       int     `json:"home_corners" validate:"min=0"`
	AwayCorners       int     `json:"away_corners" validate:"min=0"`
	HomeFouls         int     `json:"home_fouls" validate:"min=0"`
	AwayFouls         int     `json:"away_fouls" validate:"min=0"`
	HomeYellowCards   int     `json:"home_yellow_cards" validate:"min=0"`
	AwayYellowCards   int     `json:"away_yellow_cards" validate:"min=0"`
	HomeRedCards      int     `json:"home_red_cards" validate:"min=0"`
	AwayRedCards      int     `json:"away_red_cards" validate:"min=0"`
	HomeOffsides      int     `json:"home_offsides" validate:"min=0"`
	AwayOffsides      int     `json:"away_offsides" validate:"min=0"`
	HomePasses        int     `json:"home_passes" validate:"min=0"`
	AwayPasses        int     `json:"away_passes" validate:"min=0"`
	HomePassAccuracy  float64 `json:"home_pass_accuracy" validate:"min=0,max=100"`
	AwayPassAccuracy  float64 `json:"away_pass_accuracy" validate:"min=0,max=100"`
	HomeSaves         int     `json:"home_saves" validate:"min=0"`
	AwaySaves         int     `json:"away_saves" validate:"min=0"`
}

type MatchStatisticsUpdate struct {
	HomePossession    float64 `json:"home_possession" validate:"omitempty,min=0,max=100"`
	AwayPossession    float64 `json:"away_possession" validate:"omitempty,min=0,max=100"`
	HomeShots         int     `json:"home_shots" validate:"omitempty,min=0"`
	AwayShots         int     `json:"away_shots" validate:"omitempty,min=0"`
	HomeShotsOnTarget int     `json:"home_shots_on_target" validate:"omitempty,min=0"`
	AwayShotsOnTarget int     `json:"away_shots_on_target" validate:"omitempty,min=0"`
	HomeCorners       int     `json:"home_corners" validate:"omitempty,min=0"`
	AwayCorners       int     `json:"away_corners" validate:"omitempty,min=0"`
	HomeFouls         int     `json:"home_fouls" validate:"omitempty,min=0"`
	AwayFouls         int     `json:"away_fouls" validate:"omitempty,min=0"`
	HomeYellowCards   int     `json:"home_yellow_cards" validate:"omitempty,min=0"`
	AwayYellowCards   int     `json:"away_yellow_cards" validate:"omitempty,min=0"`
	HomeRedCards      int     `json:"home_red_cards" validate:"omitempty,min=0"`
	AwayRedCards      int     `json:"away_red_cards" validate:"omitempty,min=0"`
	HomeOffsides      int     `json:"home_offsides" validate:"omitempty,min=0"`
	AwayOffsides      int     `json:"away_offsides" validate:"omitempty,min=0"`
	HomePasses        int     `json:"home_passes" validate:"omitempty,min=0"`
	AwayPasses        int     `json:"away_passes" validate:"omitempty,min=0"`
	HomePassAccuracy  float64 `json:"home_pass_accuracy" validate:"omitempty,min=0,max=100"`
	AwayPassAccuracy  float64 `json:"away_pass_accuracy" validate:"omitempty,min=0,max=100"`
	HomeSaves         int     `json:"home_saves" validate:"omitempty,min=0"`
	AwaySaves         int     `json:"away_saves" validate:"omitempty,min=0"`
}

type MatchStatisticsResponse struct {
	ID                uint      `json:"id"`
	MatchID           uint      `json:"match_id"`
	HomeTeamID        uint      `json:"home_team_id"`
	AwayTeamID        uint      `json:"away_team_id"`
	HomePossession    float64   `json:"home_possession"`
	AwayPossession    float64   `json:"away_possession"`
	HomeShots         int       `json:"home_shots"`
	AwayShots         int       `json:"away_shots"`
	HomeShotsOnTarget int       `json:"home_shots_on_target"`
	AwayShotsOnTarget int       `json:"away_shots_on_target"`
	HomeCorners       int       `json:"home_corners"`
	AwayCorners       int       `json:"away_corners"`
	HomeFouls         int       `json:"home_fouls"`
	AwayFouls         int       `json:"away_fouls"`
	HomeYellowCards   int       `json:"home_yellow_cards"`
	AwayYellowCards   int       `json:"away_yellow_cards"`
	HomeRedCards      int       `json:"home_red_cards"`
	AwayRedCards      int       `json:"away_red_cards"`
	HomeOffsides      int       `json:"home_offsides"`
	AwayOffsides      int       `json:"away_offsides"`
	HomePasses        int       `json:"home_passes"`
	AwayPasses        int       `json:"away_passes"`
	HomePassAccuracy  float64   `json:"home_pass_accuracy"`
	AwayPassAccuracy  float64   `json:"away_pass_accuracy"`
	HomeSaves         int       `json:"home_saves"`
	AwaySaves         int       `json:"away_saves"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
