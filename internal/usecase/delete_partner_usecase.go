package usecase

import (
	"fmt"

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

	contracts, err := dpuc.repositories.ContractRepository.PartnerHasEnableContract(id)
	if err != nil {
		return err
	}

	if len(contracts) > 0 {
		return fmt.Errorf("impossible delete partner because it has enable contract")
	}

	return dpuc.repositories.PartnerRepository.Delete(id)
}
