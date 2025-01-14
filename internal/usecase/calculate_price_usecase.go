package usecase

import (
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/pkg/utils"
)

type CalculatePriceUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
}

func NewCalculatePriceUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
) *CalculatePriceUseCase {
	return &CalculatePriceUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
	}
}

func (cpuc *CalculatePriceUseCase) CalculatePrice(
	origin, destination string, amount float64,
) (*float64, error) {
	km, err := cpuc.adapters.AddressService.GetDistance(origin, destination)
	if err != nil {
		return nil, err
	}
	value := utils.CalculateContract(*km, amount)
	return &value, nil
}
