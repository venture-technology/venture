package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/segmentio/kafka-go"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/internal/usecase/email"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
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
		Brokers:   []string{conf.Uchiha.Address},
		Topic:     conf.Uchiha.Queue,
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	defer reader.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {

			var email *entity.Email

			message, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Fatalf("Erro ao ler mensagem do Kafka: %v", err)
			}

			email, err = email.Unserialize(&message)
			if err != nil {
				log.Fatalf("Erro ao unserializar mensagem do Kafka: %v", err)
			}
			log.Printf("Message to -->: %s", email)

			err = c.emailUseCase.SendEmail(context.Background(), email)
			if err != nil {
				log.Fatalf("Erro ao enviar email: %v", err)
			}

			err = c.emailUseCase.Record(context.Background(), email)
			if err != nil {
				log.Fatalf("Erro ao gravar record do email: %v", err)
			}
		}
	}()

	<-signals
}

func main() {
	config, err := config.Load("../../config/config.yaml")
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Cloud.Region),
		Credentials: credentials.NewStaticCredentials(config.Cloud.AccessKey, config.Cloud.SecretKey, config.Cloud.Token),
	})
	if err != nil {
		log.Fatalf("failed to create session at aws: %v", err)
	}

	clientOptions := options.Client().ApplyURI(config.Mongo.Address)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	logger, _ := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	awsRepository := repository.NewAwsRepository(sess, logger)
	emailRepository := repository.NewEmailRepository(client, config.Mongo.Database, config.Mongo.Collection, logger)

	emailUseCase := email.NewEmailUseCase(emailRepository, awsRepository, logger)

	consumer := NewConsumer(emailUseCase)
	log.Print("initing service: uchiha")
	consumer.StartConsumer()
}
