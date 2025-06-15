package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/domain/service/payments"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type CancelContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	payments     payments.Payments
}

func NewCancelContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	payments payments.Payments,
) *CancelContractUseCase {
	return &CancelContractUseCase{
		repositories: repositories,
		logger:       logger,
		payments:     payments,
	}
}

func (ccuc *CancelContractUseCase) CancelContract(ctx context.Context, uuid uuid.UUID) error {
	contract, err := ccuc.repositories.ContractRepository.GetByUUID(uuid)
	if err != nil {
		return err
	}

	err = ccuc.payments.CancelPreApproval(ctx, contract.PreApprovalID)
	if err != nil {
		return err
	}

	return ccuc.repositories.ContractRepository.Cancel(uuid)
}
