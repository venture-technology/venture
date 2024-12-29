package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/entity"
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

func (aiuc *AcceptInviteUseCase) AcceptInvite(invite *entity.Invite) error {
	err := aiuc.validatePartnership(invite)
	if err != nil {
		return err
	}
	return aiuc.repositories.InviteRepository.Create(invite)
}

func (aiuc *AcceptInviteUseCase) validatePartnership(invite *entity.Invite) error {
	partners, err := aiuc.repositories.PartnerRepository.ArePartner(invite.Driver.CNH, invite.School.CNPJ)
	if err != nil {
		return err
	}
	if partners {
		return fmt.Errorf("driver %s and school %s are already partners", invite.Driver.CNH, invite.School.CNPJ)
	}
	return nil
}
