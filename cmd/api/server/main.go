package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/venture-technology/venture/cmd/api/server/routes/v1"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/setup"
)

func main() {
	envs, err := config.Load("../../../config/config.yaml")
	if err != nil {
		envs, err = config.Load("config/config.yaml")
		if err != nil {
			panic(err)
		}
	}

	setup := setup.NewSetup()
	setup.Logger("venture-server")
	setup.Cache()
	setup.Postgres()
	setup.Repositories()
	setup.Bucket()
	setup.Email()
	setup.Adapters()

	setup.Finish()

	serverPort := envs.Server.Port
	server := setupServer(envs)
	server.Run(fmt.Sprintf(":%s", serverPort))
}

func setupServer(config *config.Config) *gin.Engine {
	router := gin.Default()
	router.GET("/status", getStatus)

	// router.POST("/hs-account-events", func(httpContext *gin.Context) {
	// 	requestBody, err := httpContext.GetRawData()
	// 	if err != nil {
	// 		httpContext.JSON(http.StatusInternalServerError, exceptions.InvalidBodyContentResponseError(err))
	// 		return
	// 	}

	// 	infra.App.Logger.Infof(fmt.Sprintf("requestParams: %s", string(requestBody)))

	// 	httpContext.Header("Content-Type", "text/plain")
	// 	httpContext.String(http.StatusOK, "Hello API Event Received")
	// })

	apisV1 := router.Group("/api/v1")
	apisV1.Use(configHeaders())
	// apisV1.Use(middleware.AuthMiddleware(config))
	v1.NewV1Controller().V1Routes(apisV1)

	return router
}

func getStatus(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("charset", "utf-8")
	c.Header("app_version", "2025.01.07 02:38")
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
