package usecase

import (
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

func (aiuc *AcceptInviteUseCase) AcceptInvite(id string) error {
	return aiuc.repositories.InviteRepository.Accept(id)
}
