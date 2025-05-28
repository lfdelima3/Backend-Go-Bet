package routes

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/lfdelima3/Backend-Go-Bet/src/controller"
	"github.com/lfdelima3/Backend-Go-Bet/src/middleware"
)

func SetupGamer(r *gin.Engine) {
	r.POST("/login", usercontroller.Login)
	r.POST("/register", usercontroller.Register)

	user := r.Group("/usuario")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/", usercontroller.ListarUsuarios)
		user.GET("/listarusuario:id", usercontroller.GetUser)
		user.PUT("/atualizaruser:id", usercontroller.UpdateUser)
		user.DELETE("/deleteuser:id", usercontroller.DeleteUser)
	}
}
