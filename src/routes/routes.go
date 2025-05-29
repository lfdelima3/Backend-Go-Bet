package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/controller"
	"github.com/lfdelima3/Backend-Go-Bet/src/middleware"
	"github.com/lfdelima3/Backend-Go-Bet/src/util"
)

func SetupRouter(r *gin.Engine, cache *util.Cache) {
	// Configuração do rate limiter
	rateLimiter := middleware.NewRateLimiter(100, time.Minute)
	r.Use(rateLimiter.RateLimit())

	// Configuração do cache middleware
	cacheMiddleware := middleware.NewCacheMiddleware(cache)

	// Rotas públicas
	r.POST("/auth/login", controller.Login)
	r.POST("/auth/register", controller.Register)

	// Grupo de rotas autenticadas
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		// Rotas de usuário
		users := auth.Group("/users")
		{
			users.GET("/", cacheMiddleware.CacheGet(5*time.Minute), controller.ListarUsuarios)
			users.GET("/:id", cacheMiddleware.CacheGetWithKey(util.UserCacheKey, 5*time.Minute), controller.GetUser)
			users.PUT("/:id", cacheMiddleware.InvalidateCache(util.UserCacheKey), controller.UpdateUser)
			users.DELETE("/:id", cacheMiddleware.InvalidateCache(util.UserCacheKey), controller.DeleteUser)
		}

		// Rotas de times
		teams := auth.Group("/teams")
		{
			teams.POST("/", cacheMiddleware.InvalidateCache(util.TeamsCacheKey), controller.CreateClube)
			teams.GET("/", cacheMiddleware.CacheGet(util.TeamCacheExpiry), controller.GetClubes)
			teams.GET("/:id", cacheMiddleware.CacheGetWithKey(util.TeamCacheKey, util.TeamCacheExpiry), controller.GetClubePorID)
			teams.PUT("/:id", cacheMiddleware.InvalidateCache(util.TeamCacheKey), controller.UpdateClube)
			teams.DELETE("/:id", cacheMiddleware.InvalidateCache(util.TeamCacheKey), controller.DeleteClube)
		}

		// Rotas de campeonatos
		tournaments := auth.Group("/tournaments")
		{
			tournaments.POST("/", cacheMiddleware.InvalidateCache(util.TournamentsCacheKey), controller.CreateCampeonato)
			tournaments.GET("/", cacheMiddleware.CacheGet(util.TournamentCacheExpiry), controller.GetCampeonatos)
			tournaments.GET("/:id", cacheMiddleware.CacheGetWithKey(util.TournamentCacheKey, util.TournamentCacheExpiry), controller.GetCampeonatosByID)
			tournaments.PUT("/:id", cacheMiddleware.InvalidateCache(util.TournamentCacheKey), controller.UpdateCampeonato)
			tournaments.DELETE("/:id", cacheMiddleware.InvalidateCache(util.TournamentCacheKey), controller.DeleteCampeonato)
		}

		// Rotas de apostas
		bets := auth.Group("/bets")
		{
			bets.POST("/", cacheMiddleware.InvalidateCache(util.UserBetsCacheKey), controller.CreateAposta)
			bets.GET("/", cacheMiddleware.CacheGetWithKey(util.UserBetsCacheKey, util.BetCacheExpiry), controller.GetApostasUsuario)
			bets.DELETE("/:id", cacheMiddleware.InvalidateCache(util.BetCacheKey), controller.DeleteAposta)
		}

		// Rotas de partidas
		matches := auth.Group("/matches")
		{
			matches.POST("/", cacheMiddleware.InvalidateCache(util.MatchesCacheKey), controller.CreatePartida)
			matches.GET("/", cacheMiddleware.CacheGet(util.MatchCacheExpiry), controller.GetPartidas)
			matches.GET("/:id", cacheMiddleware.CacheGetWithKey(util.MatchCacheKey, util.MatchCacheExpiry), controller.GetPartidaByID)
			matches.PUT("/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.UpdatePartida)
			matches.DELETE("/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.DeletePartida)
		}

		// Rotas de times em partidas
		matchTeams := auth.Group("/match-teams")
		{
			matchTeams.POST("/", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.CreatePartidaClube)
			matchTeams.GET("/", cacheMiddleware.CacheGet(util.MatchCacheExpiry), controller.GetPartidasClubes)
			matchTeams.DELETE("/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.DeletePartidaClube)
		}

		// Rotas de eventos da partida
		matchEvents := auth.Group("/match-events")
		{
			// Gols
			matchEvents.POST("/goals", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.CreateGol)
			matchEvents.GET("/goals", cacheMiddleware.CacheGet(util.MatchCacheExpiry), controller.GetGols)
			matchEvents.GET("/goals/:id", cacheMiddleware.CacheGetWithKey(util.MatchCacheKey, util.MatchCacheExpiry), controller.GetGolByID)
			matchEvents.PUT("/goals/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.UpdateGol)
			matchEvents.DELETE("/goals/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.DeleteGol)

			// Cartões
			matchEvents.POST("/cards", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.CreateCartao)
			matchEvents.GET("/cards", cacheMiddleware.CacheGet(util.MatchCacheExpiry), controller.GetCartoes)
			matchEvents.GET("/cards/:id", cacheMiddleware.CacheGetWithKey(util.MatchCacheKey, util.MatchCacheExpiry), controller.GetCartaoByID)
			matchEvents.PUT("/cards/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.UpdateCartao)
			matchEvents.DELETE("/cards/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.DeleteCartao)

			// Faltas
			matchEvents.POST("/fouls", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.CreateFalta)
			matchEvents.GET("/fouls", cacheMiddleware.CacheGet(util.MatchCacheExpiry), controller.GetFaltas)
			matchEvents.GET("/fouls/:id", cacheMiddleware.CacheGetWithKey(util.MatchCacheKey, util.MatchCacheExpiry), controller.GetFaltaByID)
			matchEvents.PUT("/fouls/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.UpdateFalta)
			matchEvents.DELETE("/fouls/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.DeleteFalta)

			// Substituições
			matchEvents.POST("/substitutions", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.CreateSubstituicao)
			matchEvents.GET("/substitutions", cacheMiddleware.CacheGet(util.MatchCacheExpiry), controller.GetSubstituicoes)
			matchEvents.GET("/substitutions/:id", cacheMiddleware.CacheGetWithKey(util.MatchCacheKey, util.MatchCacheExpiry), controller.GetSubstituicaoByID)
			matchEvents.PUT("/substitutions/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.UpdateSubstituicao)
			matchEvents.DELETE("/substitutions/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.DeleteSubstituicao)

			// Laterais
			matchEvents.POST("/throw-ins", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.CreateLateral)
			matchEvents.GET("/throw-ins", cacheMiddleware.CacheGet(util.MatchCacheExpiry), controller.GetLaterais)
			matchEvents.GET("/throw-ins/:id", cacheMiddleware.CacheGetWithKey(util.MatchCacheKey, util.MatchCacheExpiry), controller.GetLateralByID)
			matchEvents.PUT("/throw-ins/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.UpdateLateral)
			matchEvents.DELETE("/throw-ins/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.DeleteLateral)

			// Escanteios
			matchEvents.POST("/corners", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.CreateEscanteio)
			matchEvents.GET("/corners", cacheMiddleware.CacheGet(util.MatchCacheExpiry), controller.GetEscanteios)
			matchEvents.GET("/corners/:id", cacheMiddleware.CacheGetWithKey(util.MatchCacheKey, util.MatchCacheExpiry), controller.GetEscanteioByID)
			matchEvents.PUT("/corners/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.UpdateEscanteio)
			matchEvents.DELETE("/corners/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.DeleteEscanteio)
		}
	}
}
