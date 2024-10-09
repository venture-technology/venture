package infra

import (
	"database/sql"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/venture-technology/venture/cmd/api/settings"
	"github.com/venture-technology/venture/config"
	"go.uber.org/zap"
)

type Application struct {
	Config   *config.Config
	router   *gin.Engine
	V1       *gin.RouterGroup
	Adm      *gin.RouterGroup
	Database *sql.DB
	Cloud    *session.Session
	Cache    *redis.Client
	Logger   *zap.Logger
}

func NewApplication(router *gin.Engine, V1 *gin.RouterGroup, Adm *gin.RouterGroup) *Application {

	router.Use(settings.SetHeaders())

	Config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := sql.Open("postgres", postgres(Config.Database))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(Config.Cloud.Region),
		Credentials: credentials.NewStaticCredentials(Config.Cloud.AccessKey, Config.Cloud.SecretKey, Config.Cloud.Token),
	})
	if err != nil {
		log.Fatalf("failed to create session at aws: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     Config.Cache.Address,
		Password: Config.Cache.Password,
		DB:       0,
	})

	zap, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	return &Application{
		Config:   Config,
		router:   router,
		V1:       V1,
		Adm:      Adm,
		Database: db,
		Cloud:    sess,
		Cache:    rdb,
		Logger:   zap,
	}
}

func postgres(dbconfig config.Database) string {
	return "user=" + dbconfig.User +
		" password=" + dbconfig.Password +
		" dbname=" + dbconfig.Name +
		" host=" + dbconfig.Host +
		" port=" + dbconfig.Port +
		" sslmode=disable"
}
