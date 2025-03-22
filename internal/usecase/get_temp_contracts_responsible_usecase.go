package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type GetTempContractsResponsibleUseCase struct {
	Repositories *persistence.PostgresRepositories
	Logger       contracts.Logger
}

func NewGetTempContractsResponsibleUseCase(repos *persistence.PostgresRepositories, log contracts.Logger) *GetTempContractsResponsibleUseCase {
	return &GetTempContractsResponsibleUseCase{
		Repositories: repos,
		Logger:       log,
	}
}

func (gtcuc *GetTempContractsResponsibleUseCase) GetResponsibleTempContracts(cpf string) ([]value.GetTempContracts, error) {
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
		DriverCNH:             TempContract.DriverCNH,
		SchoolCNPJ:            TempContract.SchoolCNPJ,
		KidRG:                 TempContract.KidRG,
		ResponsibleCPF:        TempContract.ResponsibleCPF,
		SignatureRequestID:    TempContract.SignatureRequestID,
		UUID:                  TempContract.UUID,
		CreatedAt:             TempContract.CreatedAt,
		ExpiredAt:             TempContract.ExpiredAt,
		DriverAssignedAt:      TempContract.DriverSignedAt.Int64,
		ResponsibleAssignedAt: TempContract.ResponsibleSignedAt.Int64,
	}
}
