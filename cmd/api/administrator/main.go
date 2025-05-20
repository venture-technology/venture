package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	admin "github.com/venture-technology/venture/cmd/api/administrator/routes/admin"
	_ "github.com/venture-technology/venture/docs"
	"github.com/venture-technology/venture/internal/setup"
)

// @title Venture API
// @version 0.1
// @description Venture Backend Major API
// @termsOfService http://swagger.io/terms/
// @host localhost:9999
// @BasePath /api/admin/
func main() {
	setup := setup.NewSetup()
	setup.Logger("venture-server")
	setup.Cache()
	setup.Postgres()
	setup.Repositories()
	setup.Bucket()
	setup.Email()
	setup.Adapters()
	setup.Converters()

	setup.Finish()

	serverPort := viper.GetString("SERVER_PORT")
	server := setupServer()
	server.Run(fmt.Sprintf(":%s", serverPort))
}

func setupServer() *gin.Engine {
	router := gin.Default()
	router.GET("/status", getStatus)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apisAdmin := router.Group("/api/admin")
	apisAdmin.Use(configHeaders())
	admin.NewAdminController().AdminRoutes(apisAdmin)

	return router
}

func getStatus(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("charset", "utf-8")
	c.Header("app_version", "2025.01.07 02:38")
	c.String(http.StatusOK, "ok, 06-05-25 version! =D")
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
