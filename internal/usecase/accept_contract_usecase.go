package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/domain/service/agreements"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type AcceptContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
}

func NewAcceptContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
) *AcceptContractUseCase {
	return &AcceptContractUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
	}
}

func (ccuc *AcceptContractUseCase) AcceptContract(asras agreements.ASRASOutput) error {
	contract, err := ccuc.createStripeItems(&asras.Contract)
	if err != nil {
		return err
	}

	// Faz o parse do UUID string para uuid.UUID
	contractUUID, err := uuid.Parse(asras.Contract.UUID)
	if err != nil {
		return err
	}

	if err := ccuc.repositories.TempContractRepository.Update(
		contractUUID,
		map[string]interface{}{
			"responsible_signed_at": asras.Signatures[0].SignedAt,
			"driver_signed_at":      asras.Signatures[1].SignedAt,
		},
	); err != nil {
		return err
	}

	return ccuc.repositories.ContractRepository.Accept(contract)
}

func (ccuc *AcceptContractUseCase) createStripeItems(contract *entity.Contract) (*entity.Contract, error) {
	alreadyExists, err := ccuc.repositories.ContractRepository.ContractAlreadyExist(contract.UUID)
	if err != nil {
		return nil, err
	}

	if alreadyExists {
		return nil, fmt.Errorf("contract already exists")
	}

	responsible, err := ccuc.repositories.ResponsibleRepository.Get(contract.ResponsibleCPF)
	if err != nil {
		return nil, err
	}

	prodt, err := ccuc.adapters.PaymentsService.CreateProduct(contract)
	if err != nil {
		return nil, err
	}

	contract.StripeProductID = prodt.ID
	pr, err := ccuc.adapters.PaymentsService.CreatePrice(contract.StripeProductID, contract.Amount)
	if err != nil {
		return nil, err
	}

	contract.StripePriceID = pr.ID
	subs, err := ccuc.adapters.PaymentsService.CreateSubscription(
		responsible.CustomerId,
		contract.StripePriceID,
	)
	if err != nil {
		return nil, err
	}

	contract.StripeSubscriptionID = subs.ID
	return contract, nil
}
