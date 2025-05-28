package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	usermodel "github.com/lfdelima3/Backend-Go-Bet/src/model"
	"github.com/lfdelima3/Backend-Go-Bet/src/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	config.Connect()
	config.DB.AutoMigrate(&usermodel.User{})

	r := gin.Default()
	routes.SetupGamer(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
