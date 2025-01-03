package usecase

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type AcceptInviteUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewAcceptInviteUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *AcceptInviteUseCase {
	return &AcceptInviteUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (aiuc *AcceptInviteUseCase) AcceptInvite(uuid uuid.UUID) error {
	return aiuc.repositories.InviteRepository.Accept(uuid)
}
