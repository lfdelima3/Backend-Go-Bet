package util

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func init() {
	// Inicializa o cliente Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // sem senha
		DB:       0,  // banco de dados padrão
	})

	// Testa a conexão
	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		LogError("Erro ao conectar ao Redis", err)
	}
}

// SetCache salva um valor no cache
func SetCache(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return redisClient.Set(ctx, key, json, expiration).Err()
}

// GetCache recupera um valor do cache
func GetCache(key string, dest interface{}) error {
	ctx := context.Background()
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

// DeleteCache remove um valor do cache
func DeleteCache(key string) error {
	ctx := context.Background()
	return redisClient.Del(ctx, key).Err()
}

// Cache representa uma instância do cache Redis
type Cache struct {
	client *redis.Client
}

// Constantes para chaves de cache
const (
	UserCacheKey        = "user:%s"
	UsersCacheKey       = "users"
	TeamCacheKey        = "team:%s"
	TeamsCacheKey       = "teams"
	TournamentCacheKey  = "tournament:%s"
	TournamentsCacheKey = "tournaments"
	MatchCacheKey       = "match:%s"
	MatchesCacheKey     = "matches"
	BetCacheKey         = "bet:%s"
	UserBetsCacheKey    = "user_bets:%s"
)

// Constantes para tempos de expiração
const (
	UserCacheExpiry       = 5 * time.Minute
	TeamCacheExpiry       = 10 * time.Minute
	TournamentCacheExpiry = 15 * time.Minute
	MatchCacheExpiry      = 2 * time.Minute
	BetCacheExpiry        = 1 * time.Minute
)

// NewCache cria uma nova instância do cache
func NewCache(addr, password string, db int) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Testa a conexão
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Cache{client: client}, nil
}

// Set armazena um valor no cache
func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, json, expiration).Err()
}

// Get recupera um valor do cache
func (c *Cache) Get(ctx context.Context, key string, dest interface{}) error {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), dest)
}

// Delete remove um valor do cache
func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

// Exists verifica se uma chave existe no cache
func (c *Cache) Exists(ctx context.Context, key string) (bool, error) {
	n, err := c.client.Exists(ctx, key).Result()
	return n > 0, err
}

// Flush limpa todo o cache
func (c *Cache) Flush(ctx context.Context) error {
	return c.client.FlushDB(ctx).Err()
}

// Close fecha a conexão com o Redis
func (c *Cache) Close() error {
	return c.client.Close()
}

// Constantes para chaves de cache
const (
	// Cache de times
	TeamCacheKey    = "team:%d"
	TeamsCacheKey   = "teams:all"
	TeamCacheExpiry = 24 * time.Hour

	// Cache de campeonatos
	TournamentCacheKey    = "tournament:%d"
	TournamentsCacheKey   = "tournaments:all"
	TournamentCacheExpiry = 24 * time.Hour

	// Cache de partidas
	MatchCacheKey    = "match:%d"
	MatchesCacheKey  = "matches:all"
	MatchCacheExpiry = 1 * time.Hour

	// Cache de apostas
	BetCacheKey      = "bet:%d"
	UserBetsCacheKey = "user:%d:bets"
	BetCacheExpiry   = 30 * time.Minute

	// Cache de estatísticas
	StatsCacheKey    = "stats:%s"
	StatsCacheExpiry = 1 * time.Hour
)
