package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type ListDriverUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListDriverUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListDriverUseCase {
	return &ListDriverUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (gduc *ListDriverUseCase) ListDriver() {

}
