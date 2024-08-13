package email

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/repository"
)

type EmailUseCase struct {
	emailRepository repository.IEmailRepository
	awsRepository   repository.IAwsRepository
}

func NewEmailUseCase(emailRepository repository.IEmailRepository, awsRepository repository.IAwsRepository) *EmailUseCase {
	return &EmailUseCase{
		emailRepository: emailRepository,
		awsRepository:   awsRepository,
	}
}

func (eu *EmailUseCase) CreateRecord(ctx context.Context, recipient *entity.Email) error {
	return eu.emailRepository.CreateRecord(ctx, recipient)
}

func (eu *EmailUseCase) SendEmail(ctx context.Context, email *entity.Email) error {
	return eu.awsRepository.SendEmail(ctx, email)
}

func (eu *EmailUseCase) UnserializeJsonToEmailDto(ctx context.Context, msg *kafka.Message) (*entity.Email, error) {
	var email *entity.Email

	err := json.Unmarshal(msg.Value, &email)
	if err != nil {
		log.Fatalf("Erro ao desserializar mensagem JSON: %v", err)
		return nil, err
	}

	return email, nil
}
