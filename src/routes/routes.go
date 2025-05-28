package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/controller"
	"github.com/lfdelima3/Backend-Go-Bet/src/middleware"
)

func SetupGamer(r *gin.Engine) {
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)

	user := r.Group("/usuario")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/", controller.ListarUsuarios)
		user.GET("/listarusuario:id", controller.GetUser)
		user.PUT("/atualizaruser:id", controller.UpdateUser)
		user.DELETE("/deleteuser:id", controller.DeleteUser)
	}

	team := r.Group("/clube")
	{
		team.POST("/create", controller.CreateClube)
		team.GET("/", controller.GetClubes)
		team.GET("/:id", controller.GetClubePorID)
		team.PUT("/atualizar:id", controller.UpdateClube)
		team.DELETE("/delete:id", controller.DeleteClube)
	}
}
