package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

// Estrutura do Use Case
type GetTempContractsUseCase struct {
	Repositories *persistence.PostgresRepositories
	Logger       contracts.Logger
}

// Função para criar uma nova instância do Use Case
func NewGetTempContractsUseCase(repos *persistence.PostgresRepositories, log contracts.Logger) *GetTempContractsUseCase {
	return &GetTempContractsUseCase{
		Repositories: repos,
		Logger:       log,
	}
}

// Método para executar a lógica de negócio
func (gtcuc *GetTempContractsUseCase) GetTempContracts(cpf string) ([]value.GetTempContracts, error) {

	// Buscar os contratos pendentes no PostgresRepositories
	contracts, err := gtcuc.Repositories.TempContractRepository.FindAllByResponsible(&cpf)
	if err != nil {
		return []value.GetTempContracts{}, err
	}
	response := []value.GetTempContracts{}
	for _, contract := range contracts {
		response = append(response, buildTempContracts(&contract))
	}
	return response, nil
}

func buildTempContracts(TempContract *entity.TempContract) value.GetTempContracts {
	return value.GetTempContracts{
		ID:                    TempContract.ID,
		SigningURL:            TempContract.SigningURL,
		Status:                TempContract.Status,
		DriverCNH:             TempContract.Status,
		SchoolCNPJ:            TempContract.SchoolCNPJ,
		KidRG:                 TempContract.KidRG,
		ResponsibleCPF:        TempContract.ResponsibleCPF,
		SignatureRequestID:    TempContract.SignatureRequestID,
		UUID:                  TempContract.UUID,
		CreatedAt:             TempContract.CreatedAt,
		ExpiredAt:             TempContract.ExpiredAt,
		DriverAssignedAt:      TempContract.DriverAssignedAt.Int64,
		ResponsibleAssignedAt: TempContract.ResponsibleAssignedAt.Int64,
	}
}
