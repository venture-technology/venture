package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
)

func AuthMiddleware(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := []byte(conf.Server.Secret)

		cookie, err := c.Cookie("token")
		if err != nil || cookie == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sem cookie de sessão"})
			c.Abort()
			return
		}

		var parsedClaims jwt.MapClaims
		token, err := jwt.ParseWithClaims(cookie, &parsedClaims, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		validateClaims(middlewareClaims{
			ctx:   c,
			kind:  parsedClaims["kind"].(string),
			token: *token,
		})

		c.Set("isAuthenticated", true)
		c.Next()
	}
}

func validateClaims(middleware middlewareClaims) {
	var mappedEvents = map[string]func(ctx *gin.Context, kind string, token jwt.Token) gin.HandlerFunc{
		"driver": func(c *gin.Context, kind string, token jwt.Token) gin.HandlerFunc {
			return func(c *gin.Context) {
				claims, ok := token.Claims.(entity.ClaimsDriver)
				if !ok || !token.Valid {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
					c.Abort()
					return
				}
				c.Set("driver", claims.Driver)
				c.Next()
			}
		},
		"school": func(c *gin.Context, kind string, token jwt.Token) gin.HandlerFunc {
			return func(c *gin.Context) {
				claims, ok := token.Claims.(entity.ClaimsSchool)
				if !ok || !token.Valid {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
					c.Abort()
					return
				}
				c.Set("school", claims.School)
				c.Next()
			}
		},
		"responsible": func(c *gin.Context, kind string, token jwt.Token) gin.HandlerFunc {
			return func(c *gin.Context) {
				claims, ok := token.Claims.(entity.ClaimsResponsible)
				if !ok || !token.Valid {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
					c.Abort()
					return
				}
				c.Set("responsible", claims.Responsible)
				c.Next()
			}
		},
	}

	if event, ok := mappedEvents[middleware.kind]; ok {
		event(middleware.ctx, middleware.kind, middleware.token)
	} else {
		middleware.ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Tipo de usuário não reconhecido"})
		middleware.ctx.Abort()
	}
}

type middlewareClaims struct {
	ctx   *gin.Context
	kind  string
	token jwt.Token
}
