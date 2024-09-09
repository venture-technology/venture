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
	"github.com/venture-technology/venture/internal/usecase/contract"
	"github.com/venture-technology/venture/internal/usecase/driver"
	"github.com/venture-technology/venture/internal/usecase/invite"
	"github.com/venture-technology/venture/internal/usecase/maps"
	"github.com/venture-technology/venture/internal/usecase/partner"
	"github.com/venture-technology/venture/internal/usecase/responsible"
	"github.com/venture-technology/venture/internal/usecase/school"

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

	config, err := config.Load("config/config.yaml")
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

	g.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})

	g.Responsible()
	g.Child()
	g.School()
	g.Driver()
	g.Invite()
	g.Partner()
	g.Contract()
	g.Maps()

	log.Print("Starting Venture API")
	g.router.Run(fmt.Sprintf(":%d", config.Server.Port))

}

func (g *Gateway) Responsible() {
	handler := handler.NewResponsibleHandler(responsible.NewResponsibleUseCase(repository.NewResponsibleRepository(g.database)))
	g.group.POST("/responsible", handler.Create)
	g.group.POST("/responsible/card", handler.SaveCard)
	g.group.GET("/responsible/:cpf", handler.Get)
	g.group.PATCH("/responsible/:cpf", handler.Update)
	g.group.DELETE("/responsible/:cpf", handler.Delete)
}

func (g *Gateway) Child() {
	handler := handler.NewChildHandler(child.NewChildUseCase(repository.NewChildRepository(g.database)))
	g.group.POST("/child", handler.Create)
	g.group.GET("/child/:rg", handler.Get)
	g.group.GET("/:cpf/child", handler.FindAll)
	g.group.PATCH("/child/:rg", handler.Update)
	g.group.DELETE("/child/:rg", handler.Delete)
}

func (g *Gateway) School() {
	handler := handler.NewSchoolHandler(school.NewSchoolUseCase(repository.NewSchoolRepository(g.database)))
	g.group.POST("/school", handler.Create)
	g.group.GET("/school", handler.FindAll)
	g.group.GET("/school/:cnpj", handler.Get)
	g.group.PATCH("/school/:cnpj", handler.Update)
	g.group.DELETE("/school/:cnpj", handler.Delete)
}

func (g *Gateway) Driver() {
	handler := handler.NewDriverHandler(driver.NewDriverUseCase(repository.NewDriverRepository(g.database), repository.NewAwsRepository(g.cloud)))
	g.group.POST("/driver", handler.Create)
	g.group.GET("/driver/:cnh", handler.Get)
	g.group.PATCH("/driver/:cnh", handler.Update)
	g.group.POST("/driver/:cnh/pix", handler.SavePix)
	g.group.POST("/driver/:cnh/bank", handler.SaveBank)
	g.group.DELETE("/driver/:cnh", handler.Delete)
	g.group.GET("/driver/:cnh/gallery", handler.GetGallery)
}

func (g *Gateway) Invite() {
	handler := handler.NewInviteHandler(invite.NewInviteUseCase(repository.NewInviteRepository(g.database), repository.NewPartnerRepository(g.database)))
	g.group.POST("/invite", handler.Create)
	g.group.GET("/invite/:id", handler.Get)
	g.group.GET("/driver/invite/:cnh", handler.FindAllByCnh)
	g.group.GET("/school/invite/:cnpj", handler.FindAllByCnpj)
	g.group.PATCH("/invite/:id/accept", handler.Accept)
	g.group.DELETE("/invite/:id/decline", handler.Decline)
}

func (g *Gateway) Partner() {
	handler := handler.NewPartnerHandler(partner.NewPartnerUseCase(repository.NewPartnerRepository(g.database)))
	g.group.GET("/partner/:id", handler.Get)
	g.group.GET("/driver/partner/:cnh", handler.FindAllByCnh)
	g.group.GET("/school/partner/:cnpj", handler.FindAllByCnpj)
	g.group.DELETE("/partner/:id", handler.Delete)
}

func (g *Gateway) Contract() {
	handler := handler.NewContractHandler(contract.NewContractUseCase(repository.NewContractRepository(g.database), repository.NewChildRepository(g.database), repository.NewDriverRepository(g.database), repository.NewSchoolRepository(g.database)))
	g.group.POST("/contract", handler.Create)
	g.group.GET("/contract/:id", handler.Get)
	g.group.GET("/driver/contract", handler.FindAllByCnh)
	g.group.GET("/school/contract", handler.FindAllByCnpj)
	g.group.GET("/responsible/contract", handler.FindAllByCpf)
	g.group.PATCH("/contract/:id/cancel", handler.Cancel)
	g.group.PATCH("/webhook/contract/:id/expired", handler.Expired)
}

func (g *Gateway) Maps() {
	handler := handler.NewMapsHandler(*maps.NeWMapsUseCase(repository.NewMapsRepository(g.database)))
	g.group.POST("/maps/price", handler.CalculatePrice)
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
