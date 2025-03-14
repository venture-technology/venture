package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type GetTempContractsDriverUseCase struct {
	Repositories *persistence.PostgresRepositories
	Logger       contracts.Logger
}

func NewGetTempContractsDriverUseCase(repos *persistence.PostgresRepositories, log contracts.Logger) *GetTempContractsDriverUseCase {
	return &GetTempContractsDriverUseCase{
		Repositories: repos,
		Logger:       log,
	}
}

func (gtcduc *GetTempContractsDriverUseCase) GetDriverTempContracts(cnh string) ([]value.GetTempContracts, error) {

	contracts, err := gtcduc.Repositories.TempContractRepository.FindAllByDriver(&cnh)
	if err != nil {
		return []value.GetTempContracts{}, err
	}
	response := []value.GetTempContracts{}
	for _, contract := range contracts {
		response = append(response, buildDriverTempContracts(&contract))
	}
	return response, nil
}

func buildDriverTempContracts(TempContract *entity.TempContract) value.GetTempContracts {
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
		DriverAssignedAt:      TempContract.DriverAssignedAt.Int64,
		ResponsibleAssignedAt: TempContract.ResponsibleAssignedAt.Int64,
	}
}
