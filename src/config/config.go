package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/lfdelima3/Backend-Go-Bet/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB é a instância global do banco de dados
var DB *gorm.DB

// Config contém todas as configurações da aplicação
type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	JWT       JWTConfig
	RateLimit RateLimitConfig
}

// ServerConfig contém as configurações do servidor
type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// DatabaseConfig contém as configurações do banco de dados
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// JWTConfig contém as configurações do JWT
type JWTConfig struct {
	SecretKey string
	Duration  time.Duration
}

// RateLimitConfig contém as configurações de rate limiting
type RateLimitConfig struct {
	Requests int
	Window   time.Duration
}

// LoadConfig carrega todas as configurações da aplicação
func LoadConfig() *Config {
	config := &Config{
		Server: ServerConfig{
			Port:         getEnv("SERVER_PORT", "8080"),
			ReadTimeout:  getDurationEnv("SERVER_READ_TIMEOUT", 15*time.Second),
			WriteTimeout: getDurationEnv("SERVER_WRITE_TIMEOUT", 15*time.Second),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "api"),
			Password: getEnv("DB_PASSWORD", "123456"),
			DBName:   getEnv("DB_NAME", "betzona"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		JWT: JWTConfig{
			SecretKey: getEnv("JWT_SECRET_KEY", "AULAPAM1"),
			Duration:  getDurationEnv("JWT_DURATION", 24*time.Hour),
		},
		RateLimit: RateLimitConfig{
			Requests: getIntEnv("RATE_LIMIT_REQUESTS", 100),
			Window:   getDurationEnv("RATE_LIMIT_WINDOW", time.Minute),
		},
	}

	// Validações de segurança
	if config.JWT.SecretKey == "your-secret-key" {
		log.Println("AVISO: JWT_SECRET_KEY não está configurado, usando valor padrão inseguro")
	}

	if config.Database.Password == "postgres" {
		log.Println("AVISO: DB_PASSWORD está usando o valor padrão, considere alterá-lo em produção")
	}

	return config
}

// getEnv retorna o valor da variável de ambiente ou o valor padrão
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getIntEnv retorna o valor inteiro da variável de ambiente ou o valor padrão
func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getDurationEnv retorna o valor de duração da variável de ambiente ou o valor padrão
func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

// Connect estabelece a conexão com o banco de dados
func Connect() {
	config := LoadConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
		config.Database.SSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	// Criar as tabelas no banco de dados
	if err := DB.AutoMigrate(&model.User{}, &model.Team{}, &model.Tournament{}, &model.Match{}, &model.MatchTeam{}, &model.MatchEvent{}, &model.MatchStatistics{}, &model.Promotion{}); err != nil {
		log.Fatalf("Falha ao criar tabelas: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso")
}
