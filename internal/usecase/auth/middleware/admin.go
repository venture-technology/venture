package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		conf := config.Get()

		input := entity.Admin{
			Email:    c.Request.Header.Get("email"),
			Password: c.Request.Header.Get("password"),
		}

		admApiKey := encodeLogin(input)

		if admApiKey != conf.Admin.ApiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			c.Abort()
			return
		}

		c.Set("isAuthenticated", true)
		c.Next()
	}
}

func encoded(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func encodeLogin(login entity.Admin) string {
	key := fmt.Sprintf("%s:%s", login.Email, login.Password)
	return encoded(key)
}
