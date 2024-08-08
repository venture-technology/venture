package main

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
	"github.com/venture-technology/venture/internal/controller"
	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/internal/service"

	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()

	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := sql.Open("postgres", NewPostgres(config.Database))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Cloud.Region),
		Credentials: credentials.NewStaticCredentials(config.Cloud.AccessKey, config.Cloud.SecretKey, config.Cloud.Token),
	})
	if err != nil {
		log.Fatalf("failed to create session at aws: %v", err)
	}

	err = Migrate(db, config.Database.Schema)
	if err != nil {
		log.Fatalf("failed to execute migrations: %v", err)
	}

	InitRoutes(router, db, sess)

	router.Run(fmt.Sprintf(":%d", config.Server.Port))

}

func NewPostgres(dbconfig config.Database) string {
	return "user=" + dbconfig.User +
		" password=" + dbconfig.Password +
		" dbname=" + dbconfig.Name +
		" host=" + dbconfig.Host +
		" port=" + dbconfig.Port +
		" sslmode=disable"
}

func Migrate(db *sql.DB, filepath string) error {
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

func InitRoutes(router *gin.Engine, db *sql.DB, aws *session.Session) {

	responsibleRepository := repository.NewResponsibleRepository(db)
	responsibleService := service.NewResponsibleService(responsibleRepository)
	responsibleController := controller.NewResponsibleController(responsibleService)

	childRepository := repository.NewChildRepository(db)
	childService := service.NewChildService(childRepository)
	childController := controller.NewChildController(childService)

	schoolRepository := repository.NewSchoolRepository(db)
	schoolService := service.NewSchoolService(schoolRepository)
	schoolController := controller.NewSchoolController(schoolService)

	awsRepository := repository.NewAWSRepository(aws)
	driverRepository := repository.NewDriverRepository(db)
	driverService := service.NewDriverService(driverRepository, awsRepository)
	driverController := controller.NewDriverController(driverService)

	partnerRepository := repository.NewPartnerRepository(db)
	partnerService := service.NewPartnerService(partnerRepository)
	partnerController := controller.NewPartnerController(partnerService)

	inviteRepository := repository.NewInviteRepository(db)
	inviteService := service.NewInviteService(inviteRepository, partnerRepository)
	inviteController := controller.NewInviteController(inviteService, partnerService)

	contractRepository := repository.NewContractRepository(db)
	contractService := service.NewContractService(contractRepository)
	contractController := controller.NewContractController(contractService)

	responsibleController.RegisterRoutes(router)
	childController.RegisterRoutes(router)
	schoolController.RegisterRoutes(router)
	driverController.RegisterRoutes(router)
	partnerController.RegisterRoutes(router)
	inviteController.RegisterRoutes(router)
	contractController.RegisterRoutes(router)
}
