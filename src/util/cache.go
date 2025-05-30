package util

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
	ctx    context.Context
}

func NewCache() *Cache {
	config := config.LoadConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	ctx := context.Background()

	// Testa a conexão
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Falha ao conectar ao Redis: %v", err)
	}

	log.Println("Conexão com o Redis estabelecida com sucesso")

	return &Cache{
		client: client,
		ctx:    ctx,
	}
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(c.ctx, key, json, expiration).Err()
}

func (c *Cache) Get(key string, dest interface{}) error {
	val, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

func (c *Cache) Delete(key string) error {
	return c.client.Del(c.ctx, key).Err()
}

func (c *Cache) Clear() error {
	return c.client.FlushDB(c.ctx).Err()
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
