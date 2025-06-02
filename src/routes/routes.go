package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/controller"
	"github.com/lfdelima3/Backend-Go-Bet/src/middleware"
	"github.com/lfdelima3/Backend-Go-Bet/src/util"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, cache util.Cache, db *gorm.DB, promotionController *controller.PromotionController) {
	// Configuração do rate limiter
	rateLimiter := middleware.NewRateLimiter(100, time.Minute)
	r.Use(rateLimiter.RateLimit())

	// Configuração do cache middleware
	cacheMiddleware := middleware.NewCacheMiddleware(cache)

	// Inicialização dos controllers
	tournamentController := controller.NewTournamentController(db)
	matchEventController := controller.NewMatchEventController(db)

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
			users.GET("/", cacheMiddleware.CacheGet(5*time.Minute), controller.ListUsers)
			users.GET("/:id", cacheMiddleware.CacheGetWithKey(util.UserCacheKey, 5*time.Minute), controller.GetUser)
			users.PUT("/:id", cacheMiddleware.InvalidateCache(util.UserCacheKey), controller.UpdateUser)
			users.DELETE("/:id", cacheMiddleware.InvalidateCache(util.UserCacheKey), controller.DeleteUser)
		}

		// Rotas de times
		teams := auth.Group("/teams")
		{
			teams.POST("/", cacheMiddleware.InvalidateCache(util.TeamsCacheKey), controller.CreateTeam)
			teams.GET("/", cacheMiddleware.CacheGet(util.TeamCacheExpiry), controller.ListTeams)
			teams.GET("/:id", cacheMiddleware.CacheGetWithKey(util.TeamCacheKey, util.TeamCacheExpiry), controller.GetTeam)
			teams.PUT("/:id", cacheMiddleware.InvalidateCache(util.TeamCacheKey), controller.UpdateTeam)
			teams.DELETE("/:id", cacheMiddleware.InvalidateCache(util.TeamCacheKey), controller.DeleteTeam)
		}

		// Rotas de campeonatos
		tournaments := auth.Group("/tournaments")
		{
			tournaments.POST("/", cacheMiddleware.InvalidateCache(util.TournamentsCacheKey), tournamentController.CreateTournament)
			tournaments.GET("/", cacheMiddleware.CacheGet(util.TournamentCacheExpiry), tournamentController.ListTournaments)
			tournaments.GET("/:id", cacheMiddleware.CacheGetWithKey(util.TournamentCacheKey, util.TournamentCacheExpiry), tournamentController.GetTournament)
			tournaments.PUT("/:id", cacheMiddleware.InvalidateCache(util.TournamentCacheKey), tournamentController.UpdateTournament)
			tournaments.DELETE("/:id", cacheMiddleware.InvalidateCache(util.TournamentCacheKey), tournamentController.DeleteTournament)
		}

		// Rotas de apostas
		bets := auth.Group("/bets")
		{
			bets.POST("/", cacheMiddleware.InvalidateCache(util.UserBetsCacheKey), controller.CreateBet)
			bets.GET("/", cacheMiddleware.CacheGetWithKey(util.UserBetsCacheKey, util.BetCacheExpiry), controller.ListUserBets)
			bets.DELETE("/:id", cacheMiddleware.InvalidateCache(util.BetCacheKey), controller.CancelBet)
		}

		// Rotas de partidas
		matches := auth.Group("/matches")
		{
			matches.POST("/", cacheMiddleware.InvalidateCache(util.MatchesCacheKey), controller.CreateMatch)
			matches.GET("/", cacheMiddleware.CacheGet(util.MatchCacheExpiry), controller.ListMatches)
			matches.GET("/:id", cacheMiddleware.CacheGetWithKey(util.MatchCacheKey, util.MatchCacheExpiry), controller.GetMatch)
			matches.PUT("/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.UpdateMatch)
			matches.DELETE("/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), controller.DeleteMatch)
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
			matchEvents.POST("/", cacheMiddleware.InvalidateCache(util.MatchCacheKey), matchEventController.CreateEvent)
			matchEvents.GET("/", cacheMiddleware.CacheGet(util.MatchCacheExpiry), matchEventController.ListEvents)
			matchEvents.GET("/:id", cacheMiddleware.CacheGetWithKey(util.MatchCacheKey, util.MatchCacheExpiry), matchEventController.GetEvent)
			matchEvents.PUT("/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), matchEventController.UpdateEvent)
			matchEvents.DELETE("/:id", cacheMiddleware.InvalidateCache(util.MatchCacheKey), matchEventController.DeleteEvent)
		}

		// Rotas de promoções
		promotions := auth.Group("/promotions")
		{
			// Rotas públicas
			promotions.GET("/active", cacheMiddleware.CacheGet(5*time.Minute), promotionController.GetActivePromotions)
			promotions.GET("", cacheMiddleware.CacheGet(5*time.Minute), promotionController.ListPromotions)
			promotions.GET("/:id", cacheMiddleware.CacheGetWithKey(util.PromotionCacheKey, 5*time.Minute), promotionController.GetPromotion)

			// Rotas protegidas (requerem autenticação e permissão de admin)
			adminRoutes := promotions.Group("")
			adminRoutes.Use(middleware.AdminMiddleware())
			{
				adminRoutes.POST("", cacheMiddleware.InvalidateCache(util.PromotionsCacheKey), promotionController.CreatePromotion)
				adminRoutes.PUT("/:id", cacheMiddleware.InvalidateCache(util.PromotionCacheKey), promotionController.UpdatePromotion)
				adminRoutes.DELETE("/:id", cacheMiddleware.InvalidateCache(util.PromotionCacheKey), promotionController.DeletePromotion)
			}
		}
	}
}
