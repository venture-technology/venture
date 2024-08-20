package gateway

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/handler"
	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/internal/usecase/child"
	"github.com/venture-technology/venture/internal/usecase/responsible"

	_ "github.com/lib/pq"
)

func SetHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("charset", "utf-8")
		c.Next()
	}
}

type Gateway struct {
	router   *gin.Engine
	group    *gin.RouterGroup
	database *sql.DB
	cloud    *session.Session
}

func NewGateway(router *gin.Engine, group *gin.RouterGroup) *Gateway {
	return &Gateway{
		router: router,
		group:  group,
	}
}

func (g *Gateway) Setup() {

	config, err := config.Load("../../config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := sql.Open("postgres", postgres(config.Database))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	g.database = db

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Cloud.Region),
		Credentials: credentials.NewStaticCredentials(config.Cloud.AccessKey, config.Cloud.SecretKey, config.Cloud.Token),
	})
	if err != nil {
		log.Fatalf("failed to create session at aws: %v", err)
	}

	g.cloud = sess

	err = migrate(g.database, config.Database.Schema)
	if err != nil {
		log.Fatalf("failed to execute migrations: %v", err)
	}

	g.Responsible()
	g.Child()

	g.router.Run(fmt.Sprintf(":%d", config.Server.Port))

}

func (g *Gateway) Responsible() {
	handler := handler.NewResponsibleHandler(responsible.NewResponsibleUseCase(repository.NewResponsibleRepository(g.database)))
	g.group.POST("/responsible", handler.Create)
	g.group.GET("/responsible/:cpf", handler.Get)
	g.group.PATCH("/responsible/:cpf", handler.Update)
	g.group.DELETE("/responsible/:cpf", handler.Delete)
}

func (g *Gateway) Child() {
	handler := handler.NewChildHandler(child.NewChildUseCase(repository.NewChildRepository(g.database)))
	g.group.POST("/child", handler.Create)
	g.group.GET("/child/:rg", handler.Get)
	g.group.PATCH("/child/:rg", handler.Update)
	g.group.DELETE("/child/:rg", handler.Delete)
}

func postgres(dbconfig config.Database) string {
	return "user=" + dbconfig.User +
		" password=" + dbconfig.Password +
		" dbname=" + dbconfig.Name +
		" host=" + dbconfig.Host +
		" port=" + dbconfig.Port +
		" sslmode=disable"
}

func migrate(db *sql.DB, filepath string) error {
	schema, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}
