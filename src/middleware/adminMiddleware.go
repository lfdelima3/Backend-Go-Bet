package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
			c.Abort()
			return
		}

		userModel, ok := user.(model.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar dados do usuário"})
			c.Abort()
			return
		}

		if userModel.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acesso negado. Apenas administradores podem acessar este recurso"})
			c.Abort()
			return
		}

		c.Next()
	}
}
