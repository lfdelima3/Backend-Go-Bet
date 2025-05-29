package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lfdelima3/Backend-Go-Bet/src/util"
)

type CacheMiddleware struct {
	cache *util.Cache
}

func NewCacheMiddleware(cache *util.Cache) *CacheMiddleware {
	return &CacheMiddleware{
		cache: cache,
	}
}

// CacheGet é um middleware para cachear respostas GET
func (cm *CacheMiddleware) CacheGet(expiration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Só cacheia requisições GET
		if c.Request.Method != http.MethodGet {
			c.Next()
			return
		}

		// Gera a chave do cache baseada na URL
		cacheKey := fmt.Sprintf("cache:%s", c.Request.URL.String())

		// Tenta obter do cache
		var response interface{}
		err := cm.cache.Get(c.Request.Context(), cacheKey, &response)
		if err == nil {
			// Se encontrou no cache, retorna o valor
			c.JSON(http.StatusOK, response)
			c.Abort()
			return
		}

		// Se não encontrou no cache, continua o processamento
		c.Next()

		// Se a resposta foi bem sucedida, salva no cache
		if c.Writer.Status() == http.StatusOK {
			// Obtém a resposta do contexto
			response := c.Keys["response"]
			if response != nil {
				cm.cache.Set(c.Request.Context(), cacheKey, response, expiration)
			}
		}
	}
}

// InvalidateCache é um middleware para invalidar cache após operações de modificação
func (cm *CacheMiddleware) InvalidateCache(pattern string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Continua o processamento
		c.Next()

		// Se a operação foi bem sucedida, invalida o cache
		if c.Writer.Status() >= 200 && c.Writer.Status() < 300 {
			// Aqui você pode implementar a lógica para invalidar chaves específicas
			// Por exemplo, se for uma operação em times, invalidar o cache de times
			// cm.cache.Delete(c.Request.Context(), fmt.Sprintf(pattern, id))
		}
	}
}

// CacheGetWithKey é um middleware para cachear respostas GET com uma chave personalizada
func (cm *CacheMiddleware) CacheGetWithKey(keyPattern string, expiration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Só cacheia requisições GET
		if c.Request.Method != http.MethodGet {
			c.Next()
			return
		}

		// Gera a chave do cache
		cacheKey := fmt.Sprintf(keyPattern, c.Param("id"))

		// Tenta obter do cache
		var response interface{}
		err := cm.cache.Get(c.Request.Context(), cacheKey, &response)
		if err == nil {
			// Se encontrou no cache, retorna o valor
			c.JSON(http.StatusOK, response)
			c.Abort()
			return
		}

		// Se não encontrou no cache, continua o processamento
		c.Next()

		// Se a resposta foi bem sucedida, salva no cache
		if c.Writer.Status() == http.StatusOK {
			// Obtém a resposta do contexto
			response := c.Keys["response"]
			if response != nil {
				cm.cache.Set(c.Request.Context(), cacheKey, response, expiration)
			}
		}
	}
}
