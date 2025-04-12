package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
)

type ResponsibleMiddleware struct {
	config config.Config
}

func NewResponsibleMiddleware(
	config config.Config,
) *ResponsibleMiddleware {
	return &ResponsibleMiddleware{
		config: config,
	}
}

func (dm *ResponsibleMiddleware) Middleware() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		secret := []byte(dm.config.Server.Secret)

		cookie, err := httpContext.Cookie("token")
		if err != nil {
			httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "cookie not found"})
			httpContext.Abort()
			return
		}

		if cookie == "" {
			httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
			httpContext.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(
			cookie,
			&entity.ClaimsResponsible{},
			func(token *jwt.Token) (interface{}, error) {
				return secret, nil
			},
		)

		if err != nil {
			httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			httpContext.Abort()
			return
		}

		claims, ok := token.Claims.(*entity.ClaimsResponsible)
		if !ok || !token.Valid {
			httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			httpContext.Abort()
			return
		}

		if claims.Role != "responsible" {
			httpContext.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
			httpContext.Abort()
			return
		}

		httpContext.Set("responsible", claims.Responsible)
		httpContext.Set("isAuthenticated", true)
		httpContext.Next()
	}
}

func (dm *ResponsibleMiddleware) GetResponsibleFromMiddleware(
	httpContext *gin.Context,
) (*entity.ClaimsResponsible, error) {
	secret := []byte(dm.config.Server.Secret)

	cookie, err := httpContext.Cookie("token")
	if err != nil {
		httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "cookie not found"})
		httpContext.Abort()
		return nil, fmt.Errorf("cookie not found: %w", err)
	}

	if cookie == "" {
		httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
		httpContext.Abort()
		return nil, fmt.Errorf("token not found: %w", err)
	}

	token, err := jwt.ParseWithClaims(
		cookie,
		&entity.ClaimsResponsible{},
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		},
	)

	if err != nil {
		httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		httpContext.Abort()
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*entity.ClaimsResponsible)
	if !ok || !token.Valid {
		httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		httpContext.Abort()
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}
