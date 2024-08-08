package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/models"
)

type ContractService struct {
	contractrepository repository.IContractRepository
}

func NewContractService(contractrepository repository.IContractRepository) *ContractService {
	return &ContractService{
		contractrepository: contractrepository,
	}
}

func (cs *ContractService) CreateContract(ctx context.Context, contract *models.Contract) error {

	// validar distancia

	// calcular valor de contrato

	// criar produto

	// criar preço

	// crio inscrição

	return cs.contractrepository.CreateContract(ctx, contract)
}

// encontre contrato por cpf e receba o status via querystring
func (cs *ContractService) FindContractsByCpf(ctx context.Context, cpf, status *string) ([]models.Contract, error) {
	return cs.contractrepository.FindContractsByCpf(ctx, cpf, status)
}

func (cs *ContractService) UpdateStatusContract(ctx context.Context, record uuid.UUID, status string) error {
	return cs.contractrepository.UpdateStatusContract(ctx, record, status)
}
