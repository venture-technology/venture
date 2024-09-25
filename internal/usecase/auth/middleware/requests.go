package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RequestMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		headerApiKey := c.Request.Header.Get("x-api-key")

		if headerApiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "acesso negado"})
			c.Abort()
			return
		}

		_, err := rdb.Get(c, headerApiKey).Result()

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "chave não encontrada"})
			c.Abort()
			return
		}

		c.Next()
	}
}
