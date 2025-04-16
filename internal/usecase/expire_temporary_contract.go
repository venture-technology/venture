package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/pkg/slicecommon"
)

type ExpireTemporaryContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewExpireTemporaryContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ExpireTemporaryContractUseCase {
	return &ExpireTemporaryContractUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (etcu *ExpireTemporaryContractUseCase) ExpireTemporaryContracts() error {
	expiredContracts, err := etcu.repositories.TempContractRepository.GetExpiredContracts()
	if err != nil {
		etcu.logger.Errorf(fmt.Sprintf("Error getting expired contracts: %v", err))
		return err
	}

	var uuids []uuid.UUID
	for _, contract := range expiredContracts {
		parsedUUID, err := uuid.Parse(contract.UUID)
		if err != nil {
			etcu.logger.Errorf(fmt.Sprintf("Error parsing UUID %s: %v", contract.UUID, err))
			continue
		}
		uuids = append(uuids, parsedUUID)
	}

	batches, err := slicecommon.BatchSlice(uuids, 100)
	if err != nil {
		etcu.logger.Errorf(fmt.Sprintf("Error batching uuids: %v", err))
		return err
	}

	for _, batch := range batches {
		err := etcu.repositories.TempContractRepository.Expire(batch)
		if err != nil {
			etcu.logger.Errorf(fmt.Sprintf("Error expiring batch: %v", err))
			return err
		}
	}

	return nil
}
