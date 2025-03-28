package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
)

type SchoolMiddleware struct {
	config config.Config
}

func NewSchoolMiddleware(
	config config.Config,
) *SchoolMiddleware {
	return &SchoolMiddleware{
		config: config,
	}
}

func (dm *SchoolMiddleware) Middleware() gin.HandlerFunc {
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
			&entity.ClaimsSchool{},
			func(token *jwt.Token) (interface{}, error) {
				return secret, nil
			},
		)

		if err != nil {
			httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			httpContext.Abort()
			return
		}

		claims, ok := token.Claims.(*entity.ClaimsSchool)
		if !ok || !token.Valid {
			httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			httpContext.Abort()
			return
		}

		httpContext.Set("school", claims.School)
		httpContext.Set("isAuthenticated", true)
		httpContext.Next()
	}
}
