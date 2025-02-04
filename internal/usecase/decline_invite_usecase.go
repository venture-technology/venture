package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type DeclineInviteUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewDeclineInviteUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *DeclineInviteUseCase {
	return &DeclineInviteUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (diuc *DeclineInviteUseCase) DeclineInvite(id string) error {
	return diuc.repositories.InviteRepository.Decline(id)
}
