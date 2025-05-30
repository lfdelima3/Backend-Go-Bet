package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/controller"
	"github.com/lfdelima3/Backend-Go-Bet/src/routes"
	"github.com/lfdelima3/Backend-Go-Bet/src/util"
)

func main() {
	// Carrega as configurações
	cfg := config.LoadConfig()

	// Inicializa o banco de dados
	config.Connect()

	// Inicializa o cache
	cache := util.NewCache()

	// Configura o Gin
	router := gin.Default()

	// Configura os controllers
	promotionController := controller.NewPromotionController(config.DB)

	// Configura as rotas
	routes.SetupRouter(router, cache, config.DB, promotionController)

	// Inicia o servidor
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
