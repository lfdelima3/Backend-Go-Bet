package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/config"
	"github.com/lfdelima3/Backend-Go-Bet/src/model"
	"github.com/lfdelima3/Backend-Go-Bet/src/util"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token Ausente ou invalido"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := util.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		userID := uint(claims["user_id"].(float64))
		var user model.User
		if err := config.DB.First(&user, userID).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Usuário não encontrado"})
			return
		}

		if user.Status != "active" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Usuário inativo ou bloqueado"})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
