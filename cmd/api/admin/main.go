package admin

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/cmd/api/admin/routes/adm"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/domain/service/middleware"
	"github.com/venture-technology/venture/internal/setup"
)

func main() {
	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	setup := setup.NewSetup()
	setup.Redis()
	setup.RedisRepositories()
	setup.Logger("venture-admin-server")

	serverPort := config.Server.Port
	server := setupServer(config)
	server.Run(fmt.Sprintf(":%s", serverPort))
}

func setupServer(config *config.Config) *gin.Engine {
	router := gin.Default()
	router.GET("/status", getStatus)

	apisAdmin := router.Group("/api/admin")
	apisAdmin.Use(configHeaders())
	apisAdmin.Use(middleware.AdminMiddleware(config))
	adm.NewAdminController().AdminRoutes(apisAdmin)

	return router
}

func getStatus(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("charset", "utf-8")
	c.Header("app_version", "2025.01.03 12:21")
	c.String(http.StatusOK, "ok")
}

func configHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, email, account")
		c.Header("Cross-Origin-Embedder-Policy", "require-corp")
		c.Header("Cross-Origin-Opener-Policy", "same-origin")
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Header("charset", "utf-8")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}
