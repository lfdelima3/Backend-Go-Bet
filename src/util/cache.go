package util

import (
	"time"
)

// Cache é uma interface para operações de cache
type Cache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string, dest interface{}) error
	Delete(key string) error
	Clear() error
	Close()
}

// NewCache cria uma nova instância de cache
func NewCache() Cache {
	// Retorna um cache dummy ou baseado em memória se não usar Redis
	return &DummyCache{}
}

// DummyCache é uma implementação de cache que não faz nada
type DummyCache struct{}

func (d *DummyCache) Set(key string, value interface{}, expiration time.Duration) error {
	return nil
}

func (d *DummyCache) Get(key string, dest interface{}) error {
	return nil
}

func (d *DummyCache) Delete(key string) error {
	return nil
}

func (d *DummyCache) Clear() error {
	return nil
}

func (d *DummyCache) Close() {
	// Não faz nada
}

// Constantes para chaves de cache
const (
	UserCacheKey         = "user:"
	TeamsCacheKey        = "teams:"
	TeamCacheKey         = "team:"
	TournamentsCacheKey  = "tournaments:"
	TournamentCacheKey   = "tournament:"
	MatchesCacheKey      = "matches:"
	MatchCacheKey        = "match:"
	UserBetsCacheKey     = "user_bets:"
	BetCacheKey          = "bet:"
	PlayerCacheKey       = "player:"
	GoalCacheKey         = "goal:"
	CardCacheKey         = "card:"
	FoulCacheKey         = "foul:"
	SubstitutionCacheKey = "substitution:"
	ThrowInCacheKey      = "throw_in:"
	CornerCacheKey       = "corner:"
	PromotionsCacheKey   = "promotions:"
	PromotionCacheKey    = "promotion:"
)

// Constantes para expiração do cache
const (
	UserCacheExpiry         = 5 * time.Minute
	TeamCacheExpiry         = 10 * time.Minute
	TournamentCacheExpiry   = 15 * time.Minute
	MatchCacheExpiry        = 30 * time.Minute
	BetCacheExpiry          = 5 * time.Minute
	PlayerCacheExpiry       = 10 * time.Minute
	GoalCacheExpiry         = 5 * time.Minute
	CardCacheExpiry         = 5 * time.Minute
	FoulCacheExpiry         = 5 * time.Minute
	SubstitutionCacheExpiry = 5 * time.Minute
	ThrowInCacheExpiry      = 5 * time.Minute
	CornerCacheExpiry       = 5 * time.Minute
	PromotionCacheExpiry    = 5 * time.Minute
)
