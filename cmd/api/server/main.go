package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/venture-technology/venture/cmd/api/server/routes/v1"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/setup"
)

func main() {
	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	setup := setup.NewSetup()
	setup.Postgres()
	setup.Redis()
	setup.Repositories()
	setup.RedisRepositories()
	setup.Bucket()
	setup.Email()
	setup.Logger("venture-server")

	serverPort := config.Server.Port
	server := setupServer()
	server.Run(fmt.Sprintf(":%s", serverPort))
}

func setupServer() *gin.Engine {
	router := gin.Default()
	router.GET("/status", getStatus)

	apisV1 := router.Group("/api/v1")
	apisV1.Use(configHeaders())
	v1.NewV1Controller().V1Routes(apisV1)

	return router
}

func getStatus(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("charset", "utf-8")
	c.Header("app_version", "2024.12.26 18:42")
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
