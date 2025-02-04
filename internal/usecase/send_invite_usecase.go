package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type SendInviteUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewSendInviteUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *SendInviteUseCase {
	return &SendInviteUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (siuc *SendInviteUseCase) SendInvite(invite *entity.Invite) error {
	invite.Status = "pending"
	return siuc.repositories.InviteRepository.Create(invite)
}
