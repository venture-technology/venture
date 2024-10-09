package email

import (
	"context"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/repository"
	"go.uber.org/zap"
)

type EmailUseCase struct {
	emailRepository repository.IEmailRepository
	awsRepository   repository.IAwsRepository
	logger          *zap.Logger
}

func NewEmailUseCase(
	emailRepository repository.IEmailRepository,
	awsRepository repository.IAwsRepository,
	logger *zap.Logger,
) *EmailUseCase {
	return &EmailUseCase{
		emailRepository: emailRepository,
		awsRepository:   awsRepository,
		logger:          logger,
	}
}

func (eu *EmailUseCase) Record(ctx context.Context, email *entity.Email) error {

	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	email.ID = id

	return eu.emailRepository.Record(ctx, email)
}

func (eu *EmailUseCase) SendEmail(ctx context.Context, email *entity.Email) error {
	return eu.awsRepository.SendEmail(ctx, email)
}
