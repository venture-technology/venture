package usecase

import (
	"github.com/venture-technology/venture/internal/domain/adapter"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/pkg/utils"
)

type CalculatePriceUseCase struct {
	repositories  *persistence.PostgresRepositories
	logger        contracts.Logger
	googleAdapter adapter.IGoogleAdapter
}

func NewCalculatePriceUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	googleAdapter adapter.IGoogleAdapter,
) *CalculatePriceUseCase {
	return &CalculatePriceUseCase{
		repositories:  repositories,
		logger:        logger,
		googleAdapter: googleAdapter,
	}
}

func (cpuc *CalculatePriceUseCase) CalculatePrice(
	origin, destination string, amount float64,
) (*float64, error) {
	km, err := cpuc.googleAdapter.GetDistance(origin, destination)
	if err != nil {
		return nil, err
	}
	value := utils.CalculateContract(*km, amount)
	return &value, nil
}
