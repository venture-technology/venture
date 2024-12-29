package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type DeletePartnerUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewDeletePartnerUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *DeletePartnerUseCase {
	return &DeletePartnerUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (dpuc *DeletePartnerUseCase) DeletePartner(id string) error {
	return dpuc.repositories.PartnerRepository.Delete(id)
}
