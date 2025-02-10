package usecase

import (
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type CreateContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
}

func NewCreateContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
) *CreateContractUseCase {
	return &CreateContractUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
	}
}

func (ccuc *CreateContractUseCase) CreateContract(contract *entity.Contract) error {
	// 1. pegar template em html
	// 2. substituir variaveis do template
	// 3. chamar o converter de base64
	// 4. enviar requisição para o docusign
	return nil
}
