package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/segmentio/kafka-go"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/internal/usecase/email"
)

type consumer struct {
	emailUseCase *email.EmailUseCase
}

func NewConsumer(emailUseCase *email.EmailUseCase) *consumer {
	return &consumer{
		emailUseCase: emailUseCase,
	}
}

func (c *consumer) StartConsumer() {
	conf := config.Get()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{conf.Messaging.Brokers},
		Topic:     conf.Messaging.Topic,
		Partition: 1,
		GroupID:   "reader.kafka.group",
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	defer reader.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {

			message, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Fatalf("Erro ao ler mensagem do Kafka: %v", err)
			}

			email, err := c.emailUseCase.UnserializeJsonToEmailDto(context.Background(), &message)
			if err != nil {
				log.Fatalf("Erro ao unserializar mensagem do Kafka: %v", err)
			}
			log.Printf("Message to -->: %s", email)

			err = c.emailUseCase.SendEmail(context.Background(), email)
			if err != nil {
				log.Fatalf("Erro ao enviar email: %v", err)
			}

			err = c.emailUseCase.CreateRecord(context.Background(), email)
			if err != nil {
				log.Fatalf("Erro ao gravar record do email: %v", err)
			}

			log.Println("Venture-Microservice-Email: Message found in Queue, Email sended.")
		}
	}()

	<-signals
}

func main() {

	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = migrate(db, config.Database.Schema)
	if err != nil {
		log.Fatalf("failed to execute migrations: %v", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Cloud.Region),
		Credentials: credentials.NewStaticCredentials(config.Cloud.AccessKey, config.Cloud.SecretKey, config.Cloud.Token),
	})
	if err != nil {
		log.Fatalf("failed to create session at aws: %v", err)
	}

	awsRepository := repository.NewAwsRepository(sess)
	emailRepository := repository.NewEmailRepository(db)

	emailUseCase := email.NewEmailUseCase(emailRepository, awsRepository)

	consumer := NewConsumer(emailUseCase)
	log.Print("initing service: email-venture-service")
	consumer.StartConsumer()

}

func newPostgres(dbconfig config.Database) string {
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
